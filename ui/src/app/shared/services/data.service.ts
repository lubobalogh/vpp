import { Injectable } from '@angular/core';
import { map, switchMap, catchError } from 'rxjs/operators';
import { Observable, forkJoin, of, BehaviorSubject, Subject } from 'rxjs';
import { KubernetesService } from './kubernetes.service';
import { VppService } from './vpp.service';
import { ContivDataModel } from '../models/contiv-data-model';
import { VppArpModel } from '../models/vpp/vpp-arp-model';
import { VppInterfaceTapModel } from '../models/vpp/vpp-interface-tap-model';
import { K8sPodModel } from '../models/k8s/k8s-pod-model';
import { K8sNodeModel } from '../models/k8s/k8s-node-model';
import { VppIpamModel } from '../models/vpp/vpp-ipam-model';
import { AppConfig } from '../../app-config';

@Injectable({
  providedIn: 'root'
})
export class DataService {

  public contivData: ContivDataModel;
  public isDataLoading: Subject<boolean> = new Subject<boolean>();
  public isContivDataLoaded: BehaviorSubject<boolean> = new BehaviorSubject(false);
  public preventRefreshSubject: Subject<boolean> = new Subject<boolean>();

  constructor(
    private k8sService: KubernetesService,
    private vppService: VppService
  ) {
    this.loadData(true);
  }

  public preventRefresh() {
    this.preventRefreshSubject.next(true);
  }

  public allowRefresh() {
    this.preventRefreshSubject.next(false);
  }

  public loadData(first?: Boolean) {
    this.isDataLoading.next(true);

    if (!first) {
      this.isContivDataLoaded.next(false);
    }

    this.getPodsNamespaces().subscribe(res => {
      this.contivData = new ContivDataModel();

      res.forEach(r => {
        if (r) {
          this.contivData.addData(r);
        }
      });
      this.isContivDataLoaded.next(true);
      this.isDataLoading.next(false);
    });
  }

  public getPodsNamespaces() {
    return this.getNetworkData().pipe(
      switchMap(data => {
        const observables = data.map(d => {
          if (!d) {
            return of(null);
          }

          return this.k8sService.loadNamespaces().pipe(
            map(res => {
              const namespaces = {
                namespaces: res
              };

              return Object.assign(d, namespaces);
            })
          );
        });

        return forkJoin(observables);
      })
    );
  }

  public getNetworkData() {
    return this.getBdByNode();
  }

  private getBdByNode() {
    return this.getVrfInterfaces().pipe(
      switchMap(data => {
        const observables = data.map(d => {
          if (!d) {
            return of(null);
          }

          const url = this.getUrl(d.node.name);

          return this.vppService.getBridgeDomains(url).pipe(
            map(res => {
              const bdObj = {
                bd: res
              };

              return Object.assign(d, bdObj);
            })
          );
        });

        return forkJoin(observables);
      })
    );
  }

  private getArpByIp(nodeName: string, ip: string): Observable<VppArpModel> {
    const url = this.getUrl(nodeName);

    return this.vppService.getArps(url).pipe(
      map(res => res.find(e => e.IP === ip))
    );
  }

  private getTapInterfaceByName(nodeName: string, ifName: string): Observable<VppInterfaceTapModel> {
    const url = this.getUrl(nodeName);
    return this.vppService.getTapInterfaces(url).pipe(
      map(res => res.find(e => e.name === ifName))
    );
  }

  private getInterfaceByPod(pod: K8sPodModel, nodeName: string): Observable<VppInterfaceTapModel> {
    return this.getArpByIp(nodeName, pod.podIp).pipe(
      switchMap(arp => {
        return arp ? this.getTapInterfaceByName(nodeName, arp.interface) : of(null);
      })
    );
  }

  private getInterfacesByPods(pods: K8sPodModel[], nodeName: string): Observable<VppInterfaceTapModel[]> {
    if (!pods.length) {
      return of(null);
    }

    const observables = pods.map(p => this.getInterfaceByPod(p, nodeName));

    return forkJoin(observables);
  }

  private getVrfInterfaces() {
    return this.getPodsVppIps().pipe(
      switchMap(data => {
        const observables = data.map(d => {
          if (!d) {
            return of(null);
          }

          const url = this.getUrl(d.node.name);

          return this.vppService.getInterfaces(url).pipe(
            map(res => {
              const ifacesObject = {
                interfaces: res
              };

              return Object.assign(d, ifacesObject);
            })
          );
        });

        return forkJoin(observables);
      })
    );
  }

  private getPodsVppIps() {
    return this.getPodsByNode().pipe(
      switchMap(data => {
        const observables = data.map(d => {
          if (!d) {
            return of(null);
          }

          return this.getInterfacesByPods(d.vppPods, d.node.name).pipe(
            map(res => {
              if (res) {
                res.forEach((r, i) => {
                  if (r) {
                    d.vppPods[i].tapInterface = r.name;
                    d.vppPods[i].tapInternalInterface = r.internalName;
                  }
                });
              }
              return d;
            })
          );
        });

        return forkJoin(observables);
      })
    );
  }

  private getPodsByNode() {
    return this.getIPAMbyNode().pipe(
      switchMap(data => {
        const observables = data.map(d => {
          /* TODO Move outside map */
          return this.k8sService.loadPods().pipe(
            map(res => {
              if (d) {
                const podsObj = {
                  vppPods: res.filter(pod => pod.node === d.node.name && pod.hostIp !== pod.podIp),
                  pods: res.filter(pod => pod.node === d.node.name && pod.hostIp === pod.podIp && !pod.name.includes('vswitch')),
                  vswitch: res.find(pod => pod.node === d.node.name && pod.name.includes('vswitch'))
                };

                return Object.assign(d, podsObj);
              } else {
                return null;
              }
            })
          );
        });

        return forkJoin(observables);
      })
    );
  }

  private getIPAMbyNode(): Observable<{
    node: K8sNodeModel;
    ipam: VppIpamModel;
  }[]> {
    return this.k8sService.loadNodes().pipe(
      switchMap(nodes => {
        const observables = nodes.map(n => {
          const url = this.getUrl(n.name);

          return this.vppService.getIPAM(url).pipe(
            map(ipam => {
              return {
                node: n,
                ipam: ipam
              };
            }),
            catchError(() => of(null))
          );
        });
        return forkJoin(observables);
      })
    );
  }

  public getUrl(nodeName: string): string {
    switch (nodeName) {
      case 'k8s-master':
        return AppConfig.VPP_REST_URL_MASTER;
      case 'k8s-worker1':
        return AppConfig.VPP_REST_URL_WORKER1;
      case 'k8s-worker2':
        return AppConfig.VPP_REST_URL_WORKER2;
    }
  }

}
