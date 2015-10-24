package main

import (
	"github.com/codegangsta/cli"
	"github.com/pirsquare/askme-golang/provider"
	"github.com/pirsquare/askme-golang/utils"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "askme"
	app.Usage = "AskMe Cli"
	app.Version = utils.GetVersion()

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "fields, f",
			Usage: "Specified fields to retrieve (comma-seperated)",
		},
		cli.StringFlag{
			Name:  "delimiter, d",
			Value: " | ",
			Usage: "Delimiter (default is ' | ' with spaces)",
		},
		cli.BoolFlag{
			Name:  "omit-columns, o",
			Usage: "Omit Columns",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "gcloud",
			Usage: "(gce-zone|gce-machine-type|gce-disk-type|gce-image)",
			Subcommands: []cli.Command{
				{
					Name:  "gce-zone",
					Usage: "query GCE zone",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						gcloud := provider.NewGCloud(pvd)
						gcloud.RenderGCEZone()
					},
				},

				{
					Name:  "gce-machine-type",
					Usage: "query GCE machine type",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						gcloud := provider.NewGCloud(pvd)
						gcloud.RenderGCEMachineType()
					},
				},

				{
					Name:  "gce-disk-type",
					Usage: "query GCE disk type",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						gcloud := provider.NewGCloud(pvd)
						gcloud.RenderGCEDiskType()
					},
				},

				{
					Name:  "gce-image",
					Usage: "query GCE image",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						gcloud := provider.NewGCloud(pvd)
						gcloud.RenderGCEImage()
					},
				},
			},
		},

		{
			Name:  "do",
			Usage: "(region|dist-image|app-image|size)",
			Subcommands: []cli.Command{
				{
					Name:  "region",
					Usage: "query region",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						do := provider.NewDigitalOcean(pvd)
						do.RenderRegion()
					},
				},

				{
					Name:  "dist-image",
					Usage: "query distribution image",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						do := provider.NewDigitalOcean(pvd)
						do.RenderDistImage()
					},
				},

				{
					Name:  "app-image",
					Usage: "query application image",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						do := provider.NewDigitalOcean(pvd)
						do.RenderAppImage()
					},
				},

				{
					Name:  "size",
					Usage: "query size",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						do := provider.NewDigitalOcean(pvd)
						do.RenderSize()
					},
				},
			},
		},

		{
			Name:  "aws",
			Usage: "(ec2-region|ec2-zone|ec2-instance-type)",
			Subcommands: []cli.Command{
				{
					Name:  "ec2-region",
					Usage: "query ec2 regions",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						aws := provider.NewAWS(pvd)
						aws.RenderEC2Region()
					},
				},

				{
					Name:  "ec2-zone",
					Usage: "query ec2 zones",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						aws := provider.NewAWS(pvd)
						aws.RenderEC2Zone()
					},
				},

				{
					Name:  "ec2-instance-type",
					Usage: "query ec2 instance type",
					Flags: flags,
					Action: func(c *cli.Context) {
						pvd := provider.NewProviderWithCliContext(c)
						aws := provider.NewAWS(pvd)
						aws.RenderEC2InstanceType()
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
