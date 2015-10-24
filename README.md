# AskMe
Data files for AskMe. Changes to source data goes here.

## Use Case
Try using devops tools like docker/packer/ansible to provision a cloud server with digitalocean. [On docker's page](https://docs.docker.com/machine/drivers/digital-ocean/) it says you need to provide `image`, `region` and `size`. So you tried to google for accepted values but wait, the only way to retrieve that is to use digitalocean's api.

What if you want to use docker/packer/ansible but don't ~~want~~ have the time to learn digitalocean's api? Try [AskMe](https://github.com/pirsquare/askme)

## Client Libraries
[NodeJS](https://github.com/pirsquare/askme-node)

## Supported Cloud Providers
- AWS
- Google Cloud
- Digitalocean
- Azure (Planned)
- Rackspace (Planned)
- Softlayer (Planned)

## Supported fields:
- id (Id)
- desc (Description)

## AWS
List of supported records to query:
- ec2-region
- ec2-zone
- ec2-instance-type


## Google Cloud
List of supported records to query:
- gce-zone
- gce-machine-type
- gce-disk-type
- gce-image


## Digitalocean
List of supported records to query:
- region
- dist-image (Distribution Image)
- app-image (Application Image)
- size (Instance Size)

## Note
Git subtree is used to push source data to client libraries. If you want submit PR for changes in source data, submit it here, instead of the client libraries.
