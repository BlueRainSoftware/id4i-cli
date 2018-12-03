This is the ID4i CLI README describing the usage via docker. 

For the binary distribution and development README, please refer to 
https://github.com/BlueRainSoftware/id4i-cli/blob/master/README.adoc

---

# ID4i CLI Docker Image

ID4i provides globally unique IDs for single workpieces and a platform to manage, and exchange data bound to IDs in an inter-organizational manner. For details, please refer to http://id4i.de.

You can use this the ID4i command line interface to ...

* perform manual tasks in ID4i
* include ID4i tasks in arbitrary shell scripts
* query ID4i manually or as part of existing workflows.

For a detailed documentation of ID4i see http://backend.id4i.de/docs/reference/en/reference.html#_tutorials


You can use the ID4i CLI via Docker so you don't require installing the binaries. Running `docker run bluerainsoftware/id4i-cli` displays help.
 
To configure the image with your API key and organization, you need to pass environment variables 
to the image or mount a volume containing your configuration file.

For additional details on prerequisites, please refer to 
https://github.com/BlueRainSoftware/id4i-cli/blob/master/README.adoc#usage .

## Configuration using environment variables

The typical case is that you set your configuration using environment
variables in the container and don't use a configuration file. 
If you do have a local configuration file, see the following section.

The minimum required configuration consists of the `apikey` and `secret` properties. 
Most operations also require `organization` to be set.

An example call setting all three properties looks like this:

```bash
docker run \
    -e ID4I_APIKEY=<your API key ID> \
    -e ID4I_SECRET=<API key secret> \
    -e ID4I_ORGANIZATION=<organization namespace> \
    bluerainsoftware/id4i-cli \
    info
```

The following variables are supported for configuration

* `ID4I_CONFIG` - configuration file location
* `ID4I_BACKEND` - ID4i backend to use, either `sandbox.id4i.de` or `backend.id4i.de`
* `ID4I_APIKEY` - ID of the API key used to connect
* `ID4I_SECRET` - API key secret for signing
* `ID4I_ORGANIZATION` - Organization namespace to work in

## Configuration using a config file

Assuming you have a configuration file called `.id4i.properties` in `/Users/me/.id4i/`, you can
tell `id4i` to use it by telling it the mounted file's location via command line parameter or environment
variable

```bash
docker run \
    -v /Users/me/.id4i:/home/.id4i \ 
    bluerainsoftware/id4i-cli info \
    --config /home/.id4i/.id4i.properties
```

To omit the `--config` parameter and use the default config file path, you can mount only the 
config file to the default config file location, e.g.

```bash
docker run \
    -v /Users/me/.id4i.properties:/home/.id4i.properties \ 
    bluerainsoftware/id4i-cli info
```

As an alternative, you can use the `ID4I_CONFIG` environment variable to specify the location of the mounted config file.
```bash
docker run \
    -v /Users/me/.id4i:/home/.id4i \ 
    -e ID4I_CONFIG=/home/.id4i/.id4i.properties \
    bluerainsoftware/id4i-cli info 
```

## Development

If you need additional operations to be supported, head over to the GitHub repository and see 
the README.adoc there. It has loads of additional information.
