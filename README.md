batterystats
============

A tool that collects battery statistics from your laptop and writes them to a file.
It can write to a file in JSON format, one line per battery per collection.

Installation
------------

    go install github.com/solarnz/batterystats/...

or from the [AUR](https://aur.archlinux.org/packages/batterystats-git/).

Sample output
-------------

    {"Timestamp":"2015-09-27T20:46:34.992346197+10:00","BatteryID":"BAT0","CycleCount":0,"EnergyFull":22030000,"EnergyFullDesign":23200000,"EnergyNow":17520000,"Manufacturer":"SONY","ModelName":"45N1111","PowerNow":0,"SerialNumber":"15114","Status":"Unknown","Technology":"Li-poly"}
    {"Timestamp":"2015-09-27T20:46:34.99238859+10:00","BatteryID":"BAT1","CycleCount":0,"EnergyFull":65800000,"EnergyFullDesign":71100000,"EnergyNow":59170000,"Manufacturer":"LGC","ModelName":"45N1738","PowerNow":1717000,"SerialNumber":"  331","Status":"Charging","Technology":"Li-ion"}
    {"Timestamp":"2015-09-27T20:49:45.974225483+10:00","BatteryID":"BAT0","CycleCount":0,"EnergyFull":22030000,"EnergyFullDesign":23200000,"EnergyNow":17520000,"Manufacturer":"SONY","ModelName":"45N1111","PowerNow":0,"SerialNumber":"15114","Status":"Unknown","Technology":"Li-poly"}
    {"Timestamp":"2015-09-27T20:49:45.975398187+10:00","BatteryID":"BAT1","CycleCount":0,"EnergyFull":65800000,"EnergyFullDesign":71100000,"EnergyNow":59260000,"Manufacturer":"LGC","ModelName":"45N1738","PowerNow":1704000,"SerialNumber":"  331","Status":"Charging","Technology":"Li-ion"}
    
Usage
-----

    Usage of batterystats:
      -one
        	Only collect battery statistics once
      -out string
        	Where to write the output
      -period int
        	How often in minutes to collect battery statistics (default 10)

There are also some systemd service files in the root of the repository. Install and enable these to automatically collect data to /var/log/battery.log every 10 minutes and whenever you suspend or resume your laptop.
