- [Scale Out Http Server](#scale_out_http_server)

## Scale Out Http Server <a name="scale_out_http_server"></a>

HttpServer is the easiest object to scale out, to choose the requets threshold for each instace http server, you need to
pay attention: </br>

1) What is the job of the services. </br>
2) The qps threshold must ensure that the service operates at a stable level, usually around 70% cpu load. </br>
3) Pay attention to set limit file and limit resource. <br>
4) Lightweight languages like go, rust, c++ will give very high load threshold for the same hardware. </br>
   There is a specific description of this at the
   link: https://github.com/Nghiait123456/InfraSREDevopsBackendForBigProject/tree/master/InfraSREDevops/1_LoadBalancing </br>