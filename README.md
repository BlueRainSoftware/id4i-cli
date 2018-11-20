This is the ID4i CLI README describing the usage via docker. 

For the binary distribution and development README, please refer to 
https://github.com/BlueRainSoftware/id4i-cli/blob/master/README.adoc

---

# ID4i CLI Docker Image

You can use the ID4i CLI via Docker so you don't require installing the binaries.

Running `docker run bluerainsoftware/id4i-cli:develop` displays the command help.
 
To configure the image with your API key and organization, you need to pass environment variables 
to the image or mount a volume containing your configuration file.

For additional details on prerequisites, please refer to 
https://github.com/BlueRainSoftware/id4i-cli/blob/master/README.adoc#usage .

## Configuration using environment variables


## Configuration using a config file

Assuming you have a configuration file called `.id4i.properties` in `/Users/me/.id4i/`, you can
tell `id4i` to use it by telling it the mounted file's location via command line parameter or environment
variable


```bash
docker run \
    -v /Users/w.werner/.id4i:/home/.id4i \ 
    bluerainsoftware/id4i-cli:develop info \
    --config /home/.id4i/.id4i.properties
```

To omit the `--config` parameter and use the default config file path, you can mount only the config file, e.g.

```bash
docker run \
    -v /Users/w.werner/.id4i.properties:/home/.id4i.properties \ 
    bluerainsoftware/id4i-cli:develop info
```

As an alternative, you can use the `ID4I_CONFIG` environment variable to specify the location of the mounted config file.
```bash
docker run \
    -v /Users/w.werner/.id4i:/home/.id4i \ 
    -e ID4I_CONFIG=/home/.id4i/.id4i.properties \
    bluerainsoftware/id4i-cli:develop info 
```



docker run -v /Users/w.werner/.id4i:/home/.id4i -e ID4I_CONFIG=/home/.id4i/.id4i.properties bluerainsoftware/id4i-cli:develop info 
