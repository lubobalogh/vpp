# Contiv VPP UI IPv6

Project page: https://contivpp.io/

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 7.0.2.

## Start App

### Prerequisities

- Google Chrome - running with [disabled web security](https://stackoverflow.com/questions/3102819/disable-same-origin-policy-in-chrome) (Important for accessing the APIs).
    ##### -Command in OSX:
    Open -n -a "Google Chrome" --args --disable-web-security --user-data-dir=/tmp/chrome http://localhost:4300
    ##### -Command in Windows:
    In the "Run" app, enter: chrome.exe --user-data-dir="C://Chrome dev session" --disable-web-security http://localhost:4300
    ##### -Command in Linux:
    google-chrome --user-data-dir="/var/tmp/Chrome" --disable-web-security "http://localhost:4300"
- Vagrant
- VirtualBox
- Laptop w/ 16GB of memory
- Only runs on a laptop at this time. It does not support nor can it be deployed on a remote server.
- [Postman](https://www.getpostman.com/) to look over [APIs](https://github.com/ligato/vpp-agent/tree/dev/plugins/restv2)

### Running

Shutdown all extraneous apps. You need as much memory as possible on your laptop to run the k8s cluster.

1. clone this repository (`git clone https://github.com/lubobalogh/vpp.git`)
2. change directory to vpp (`cd vpp`)
3. checkout 'ipv6' branch (`git checkout ipv6`).
4. cd into vpp/vagrant directory (`cd vagrant`)
5. from this directory run one of these commands:
- `./vagrant-start` for fully automated setup - run 2 workers in production enviroment with IPv6 enabled

You should then see this:

	CHMETZ-M-72TZ:vpp chrismetz$ cd vagrant
	CHMETZ-M-72TZ:vagrant chrismetz$ ./vagrant-start
	Starting Contiv VPP ...

	Creating a production environment, with ipv6 and 2 worker node(s)

	Creating VirtualBox DHCP server...
	Bringing machine 'k8s-gateway' up with 'virtualbox' provider...
	Bringing machine 'k8s-master' up with 'virtualbox' provider...
	Bringing machine 'k8s-worker1' up with 'virtualbox' provider...
	Bringing machine 'k8s-worker2' up with 'virtualbox' provider...
	==> k8s-gateway: Cloning VM...
	etc ... etc ...

Get a cup of coffee. It will be a few minutes.

When everything is deployed, navigate to `http://localhost:4300/` on Chrome with disabled web security - it should be automatically open if installed. If for some reason it does not automatically open, go to the CLI and type in one of the disable web security commands. This is the one for a mac:

	Open -n -a "Google Chrome" --args --disable-web-security --user-data-dir=/tmp/chrome http://localhost:4300

We need to wait a few moments (could be a couple of minutes) until k8s has completed its work and all systems are ready to go.


### Accessing APIs

#### Postman collection
[Collection](../docs/ContivVPP.postman_collection.json)

### Use 'vagrant ssh' to control nodes
After successful build, you can access created Kubernetes nodes CLI.

1. cd into vpp/vagrant (`cd /vpp/vagrant`)
2. from this directory run one of these commands:
- `vagrant ssh k8s-master` for accessing k8s-master node
- `vagrant ssh k8s-worker1` for accessing k8s-worker1 node
- `vagrant ssh k8s-worker2` for accessing k8s-worker2 node

For more details click [here](https://github.com/contiv/vpp/tree/master/vagrant).

### Shutdown App
When you finished the work with the application, run `./vagrant-shutdown` from `vpp/vagrant` folder - this will shutdown all Vagrant instances.

### Wiping all data/settings

Run `./vagrant-cleanup` from `vpp/vagrant` directory for clearing all data and settings (destroy each Vagrant instance). After this command you need to follow [running steps](#running) from the 3rd point and build application again.
