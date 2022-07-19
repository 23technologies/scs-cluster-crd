This is a reformatted version of the param-discussion (04.07.2022)
The parameters have just been sorted by type.

# Params that are required (Type 4)
KUBERNETES_VERSION: (4) Support 1.22.6 as well as  1.22.x syntax with x meaning latest (at cluster creation time)
CONTROL_PLANE_MACHINE_COUNT: (4) CONTROL_PLANE.MACHINE_FLAVOR
WORKER_MACHINE_COUNT: (4) Per machine deployment WORKER.MACHINE_COUNT

# Infra specific settings (Type 3)
OPENSTACK_CLOUD: (3) -> cloudInfrastructureReference (references clouds.yaml/secure.yaml from user in OpenStack)
OPENSTACK_FAILURE_DOMAIN: Region/Location/AZ/Zone: (3) 
OPENSTACK_CONTROL_PLANE_MACHINE_FLAVOR: (3) CONTROL_PLANE.MACHINE_FLAVOR (may be ignored by implementations without control plane on separate machines/VMs)
OPENSTACK_NODE_MACHINE_FLAVOR: (3) Per Machine deployment WORKER.MACHINE_FLAVOR

# Implementation Details (Type 2)
OPENSTACK_EXTERNAL_NETWORK_ID: (2) -> Should be auto-discovered (internet) -- only trouble is that on some clouds there are more than one external networks -> (7) advanced config
OPENSTACK_DNS_NAMESERVERS: (2) Provider should have sane defaults (either in subnet default config or in created subnet dhcp or in kubeadm)
Need to support multiple zones with one cluster => Move parameter into a per MachineDeployment parameter (and an array for the control plane)
MTU_VALUE: (2) AUTODETECT! Specific cluster-stack for Jumbo frames ...  not available on all providers. (Fallback options?)
USE_CILIUM: (2) -> could create separate cluster stacks
OPENSTACK_IMAGE_NAME: (2)
OPENSTACK_IMAGE_RAW: (2) cloud provider to chose
OPENSTACK_IMAGE_REGISTATION_EXTRA_FLAGS: (2)
OPENSTACK_CONTROL_PLANE_IP: (2)
OPENSTACK_SSH_KEY_NAME: (2)
CONTROL_PLANE_MACHINE_GEN: (2) Auto-generate 
WORKER_MACHINE_GEN: (2) Auto-generate
OPENSTACK_SRVGRP_CONTROLLER: (2)
OPENSTACK_SRVGRP_WORKER: (2)
DEPLOY_OCCM: (2) - always on
DEPLOY_CINDERCSI: (2)
ETCD_PRIO_BOOST: (2) Provider defines it such that it works
ETCD_UNSAFE_FS: (2) dito

# Opt-In Services (Provider has to offer them) (Type 5)
DEPLOY_METRICS: (5) Default on
DEPLOY_NGINX_INGRESS: (5) Default off?

# Opt-In Services (Provider can offer them) (Type 6)
DEPLOY_CERT_MANAGER: (6)
DEPLOY_FLUX: (6)
DEPLOY_HARBOR: (6) in the future

# Advanced config (Type 7)
NODE_CIDR: (7) - might want to use IPv6
