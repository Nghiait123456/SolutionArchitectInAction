- [Process](#process)


## Process <a name="process"></a>
+) Run in local: </br>
1) Create images, compose  </br>
2) Connect other service: cache, DB,.... </br>
3) Run prepare require: migrate DB, setup,... </br>
4) Run app in local </br>



+) Auto build & push docker image to AWS ECR with Github Actions: </br>
1) Create images, build images docker </br>
2) Have AWS account </br>
3) Create a repository store all images build </br>
4) Log in to AWS ECR </br>
5) Config AWS credentials </br>
6) Create some credentials allow github to access aws account </br>
7) Add user name for CICD </br>
8) Set permission for user in aws allow CICD </br>
9) Edit to gtap secret </br>
10) Set up github secret save secret for CICD </br>
11) Set up docker build latest images </br>
12) Check aws console confirm done </br>
13) Create prepare require product: DB, Cache,......
14) Store secret key of project in aws secret and auto retrieve

+) CD with EKS: </br>
1) Create new eks cluster and add worker node </br>
2) Setup networking properties for cluster </br>
3) Configure access to the Kubernetes Api Server </br>
4) Config logging </br>
5) Review and create </br>
6) Add worker node to cluster </br>
7) Config 3 options for group scaling </br>
8) Specify networking </br>

+) Control EKS: </br>
1) Use k8s cmd tool </br>
2) Verify cube control config after control </br>
3) Fetch aws eks cluster's informations </br>
4) Add permission </br>
5) Update cube config </br>
6) Check current user identity after use <br>
7) Create and give github user access to eks cluster </br>
8) Config aws eks cli use github credentials </br>
9) Use k9s for control EKS simple </br>

+) Setup K8s in EKS: </br>
1) Create a deployment  </br>
2) Config metadata </br>
3) Set specification of the deployment object </br>
4) Set up pod selected rule to control traffic </br>
5) Check resource of node </br>

+) Domain setup: </br>
1) Register a domain </br>
2) Check domain status </br>
3) Create A record </br>
4) Setup LB requets from domain to LB </br>

+) Use Ingress LB: </br>
1) config Ingress </br>
2) Deploy the Ingress </br>
3) Install Nginx Ingress Controller </br>
4) Solution LB for many cluster k8s </br>

+) Setup ssl for project: </br>
1) Setup ssl for domain </br>
2) Verify ssl </br>

+) Auto deploy K8s with Github Action: </br>
1) Setup kuber config to cluster products </br>
2) Push change to github </br>
3) Config deploy workflow </br>
4) Use K9s console verify correct images deployment </br>

+) Setup logging, monitoring: </br>
1) Setup </br>
2) Check </br>






