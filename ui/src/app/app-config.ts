export class AppConfig {
  public static K8S_REST_URL = 'http://localhost:8080/';
  public static VPP_REST_URL_MASTER = 'http://localhost:9999/';
  public static VPP_REST_URL_WORKER1 = 'http://localhost:9991/';
  public static VPP_REST_URL_WORKER2 = 'http://localhost:9992/';
  public static API_V1 = 'api/v1/';
  public static API_V1_NETWORKING = 'apis/networking.k8s.io/v1/';
  public static API_V1_CONTIV = 'contiv/v1/';
  public static API_V1_VPP = 'vpp/dump/v1/';
  public static API_V2_VPP = 'dump/vpp/v2/';
  public static ENABLE_DATA_REFRESH = true;
  public static POLLING_FREQ = 10;
}
