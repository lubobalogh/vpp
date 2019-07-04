import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { CoreService } from './core.service';
import { AppConfig } from 'src/app/app-config';
import { VppIpamModel } from '../models/vpp/vpp-ipam-model';
import { VppInterfaceModel } from '../models/vpp/vpp-interface-model';
import { VppRouteModel } from '../models/vpp/vpp-route-model';
import { VppArpModel } from '../models/vpp/vpp-arp-model';
import { VppInterfaceVxlanModel } from '../models/vpp/vpp-interface-vxlan-model';
import { VppBdModel } from '../models/vpp/vpp-bd-model';
import { VppInterfaceTapModel } from '../models/vpp/vpp-interface-tap-model';

@Injectable({
  providedIn: 'root'
})
export class VppService {

  constructor(
    private http: HttpClient,
    private coreService: CoreService
  ) { }

  public getIPAM(url: string): Observable<VppIpamModel> {
    return this.http.get(url + AppConfig.API_V1_CONTIV + 'ipam').pipe(
      map(res => {
        return this.coreService.extractObjectData(res, VppIpamModel);
      })
    );
  }

  public getIpamRaw(url: string) {
    return this.http.get(url + AppConfig.API_V1_CONTIV + 'ipam');
  }

  public getRoutes(url: string): Observable<VppRouteModel[]> {
    return this.http.get(url + AppConfig.API_V1_VPP + 'routes').pipe(
      map(res => {
        return this.coreService.extractListData(res as Array<any>, VppRouteModel);
      })
    );
  }

  public getRoutesRaw(url: string) {
    return this.http.get(url + AppConfig.API_V1_VPP + 'routes');
  }

  public getArps(url: string): Observable<VppArpModel[]> {
    return this.http.get(url + AppConfig.API_V2_VPP + 'arps').pipe(
      map(res => {
        return this.coreService.extractListData(res as Array<any>, VppArpModel);
      })
    );
  }

  public getArpsRaw(url: string) {
    return this.http.get(url + AppConfig.API_V2_VPP + 'arps');
  }

  public getInterfaces(url: string): Observable<VppInterfaceModel[]> {
    return this.http.get(url + AppConfig.API_V2_VPP + 'interfaces').pipe(
      map(res => {
        return this.coreService.extractObjectDataToArray(res, VppInterfaceModel);
      })
    );
  }

  public getInterfacesRaw(url: string) {
    return this.http.get(url + AppConfig.API_V2_VPP + 'interfaces');
  }

  public getVxlanInterfaces(url: string): Observable<VppInterfaceVxlanModel[]> {
    return this.http.get(url + AppConfig.API_V2_VPP + 'interfaces/vxlan').pipe(
      map(res => {
        return this.coreService.extractObjectDataToArray(res, VppInterfaceVxlanModel);
      })
    );
  }

  public getVxlanInterfacesRaw(url: string) {
    return this.http.get(url + AppConfig.API_V2_VPP + 'interfaces/vxlan');
  }

  public getTapInterfaces(url: string): Observable<VppInterfaceTapModel[]> {
    return this.http.get(url + AppConfig.API_V2_VPP + 'interfaces/tap').pipe(
      map(res => {
        return this.coreService.extractObjectDataToArray(res, VppInterfaceTapModel);
      })
    );
  }

  public getTapInterfacesRaw(url: string) {
    return this.http.get(url + AppConfig.API_V2_VPP + 'interfaces/tap');
  }

  public getBridgeDomains(url: string): Observable<VppBdModel[]> {
    return this.http.get(url + AppConfig.API_V2_VPP + 'bd').pipe(
      map(res => {
        return this.coreService.extractObjectDataToArray(res, VppBdModel);
      })
    );
  }

  public getbridgeDomainsRaw(url: string) {
    return this.http.get(url + AppConfig.API_V2_VPP + 'bd');
  }

  public getNatRaw(url: string) {
    return this.http.get(url + AppConfig.API_V2_VPP + 'nat/dnat');
  }

  public getVersion(url: string) {
    const headers = new HttpHeaders({
      'Content-Type':  'application/json',
    });

    const body = {'vppclicommand': 'show version'};

    return this.http.post(url + 'vpp/command', body, {headers, responseType: 'text'}).pipe(

    );
  }
}
