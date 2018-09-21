# gomountinfo

[![Build Status](https://travis-ci.org/mrostecki/gomountinfo.svg?branch=master)](https://travis-ci.org/mrostecki/gomountinfo)

Go library for parsing information from /proc/self/mountinfo

## Motivation

There are many projects written in Go which make use of mountinfo. They are
implementing mountinfo parsers on their own. The purpose of this library is to
provide an universal way of getting mount information for Go programs.

## Credits to the previous implementations

This library is **not** written from scratch. It's heavily based on previously
existing implementations of mountinfo parser in:

* [Cilium](https://github.com/cilium/cilium)
* [Kubernetes](https://github.com/kubernetes/kubernetes)
* [Moby](https://github.com/moby/moby)

All those projects are licensed under Apache License 2.0 and authors of their
mountinfo modules are credited
[here](https://github.com/mrostecki/gomountinfo/blob/master/AUTHORS).

## Example

To build the example program, please use:

```
make example
```

Then you can use the binary in `out/` directory:

```
$ ./out/example
Mount ID   Parent ID   Major   Minor   Root   Mount point                       Mount options                     Optional fields   Filesystem type   Mount source              Super options
21         68          0       20      /      /sys                              rw,nosuid,nodev,noexec,relatime   [shared:2]        sysfs             sysfs                     [rw]
22         68          0       4       /      /proc                             rw,nosuid,nodev,noexec,relatime   [shared:23]       proc              proc                      [rw]
23         68          0       6       /      /dev                              rw,nosuid                         [shared:19]       devtmpfs          devtmpfs                  [rw size=8023836k nr_inodes=2005959 mode=755]
24         21          0       7       /      /sys/kernel/security              rw,nosuid,nodev,noexec,relatime   [shared:3]        securityfs        securityfs                [rw]
25         23          0       21      /      /dev/shm                          rw,nosuid,nodev                   [shared:20]       tmpfs             tmpfs                     [rw]
26         23          0       22      /      /dev/pts                          rw,nosuid,noexec,relatime         [shared:21]       devpts            devpts                    [rw gid=5 mode=620 ptmxmode=000]
27         68          0       23      /      /run                              rw,nosuid,nodev                   [shared:22]       tmpfs             tmpfs                     [rw mode=755]
28         21          0       24      /      /sys/fs/cgroup                    ro,nosuid,nodev,noexec            [shared:4]        tmpfs             tmpfs                     [ro mode=755]
29         28          0       25      /      /sys/fs/cgroup/unified            rw,nosuid,nodev,noexec,relatime   [shared:5]        cgroup2           cgroup                    [rw nsdelegate]
30         28          0       26      /      /sys/fs/cgroup/systemd            rw,nosuid,nodev,noexec,relatime   [shared:6]        cgroup            cgroup                    [rw xattr name=systemd]
31         21          0       27      /      /sys/fs/pstore                    rw,nosuid,nodev,noexec,relatime   [shared:18]       pstore            pstore                    [rw]
32         28          0       28      /      /sys/fs/cgroup/cpuset             rw,nosuid,nodev,noexec,relatime   [shared:7]        cgroup            cgroup                    [rw cpuset]
33         28          0       29      /      /sys/fs/cgroup/net_cls,net_prio   rw,nosuid,nodev,noexec,relatime   [shared:8]        cgroup            cgroup                    [rw net_cls net_prio]
34         28          0       30      /      /sys/fs/cgroup/blkio              rw,nosuid,nodev,noexec,relatime   [shared:9]        cgroup            cgroup                    [rw blkio]
35         28          0       31      /      /sys/fs/cgroup/devices            rw,nosuid,nodev,noexec,relatime   [shared:10]       cgroup            cgroup                    [rw devices]
36         28          0       32      /      /sys/fs/cgroup/cpu,cpuacct        rw,nosuid,nodev,noexec,relatime   [shared:11]       cgroup            cgroup                    [rw cpu cpuacct]
37         28          0       33      /      /sys/fs/cgroup/hugetlb            rw,nosuid,nodev,noexec,relatime   [shared:12]       cgroup            cgroup                    [rw hugetlb]
38         28          0       34      /      /sys/fs/cgroup/freezer            rw,nosuid,nodev,noexec,relatime   [shared:13]       cgroup            cgroup                    [rw freezer]
39         28          0       35      /      /sys/fs/cgroup/memory             rw,nosuid,nodev,noexec,relatime   [shared:14]       cgroup            cgroup                    [rw memory]
40         28          0       36      /      /sys/fs/cgroup/pids               rw,nosuid,nodev,noexec,relatime   [shared:15]       cgroup            cgroup                    [rw pids]
41         28          0       37      /      /sys/fs/cgroup/perf_event         rw,nosuid,nodev,noexec,relatime   [shared:16]       cgroup            cgroup                    [rw perf_event]
42         28          0       38      /      /sys/fs/cgroup/rdma               rw,nosuid,nodev,noexec,relatime   [shared:17]       cgroup            cgroup                    [rw rdma]
68         0           254     1       /      /                                 rw,relatime                       [shared:1]        xfs               /dev/mapper/system-root   [rw attr2 inode64 noquota]
43         22          0       41      /      /proc/sys/fs/binfmt_misc          rw,relatime                       [shared:24]       autofs            systemd-1                 [rw fd=34 pgrp=1 timeout=0 minproto=5 maxproto=5 direct pipe_ino=17732]
44         23          0       42      /      /dev/hugepages                    rw,relatime                       [shared:25]       hugetlbfs         hugetlbfs                 [rw pagesize=2M]
45         21          0       8       /      /sys/kernel/debug                 rw,relatime                       [shared:26]       debugfs           debugfs                   [rw]
46         23          0       19      /      /dev/mqueue                       rw,relatime                       [shared:27]       mqueue            mqueue                    [rw]
78         68          8       1       /      /boot                             rw,relatime                       [shared:28]       ext4              /dev/sda1                 [rw stripe=4]
80         68          254     3       /      /home                             rw,relatime                       [shared:29]       xfs               /dev/mapper/system-home   [rw attr2 inode64 noquota]
382        27          0       47      /      /run/user/1000                    rw,nosuid,nodev,relatime          [shared:294]      tmpfs             tmpfs                     [rw size=1606744k mode=700 uid=1000 gid=100]
439        382         0       48      /      /run/user/1000/gvfs               rw,nosuid,nodev,relatime          [shared:348]      fuse.gvfsd-fuse   gvfsd-fuse                [rw user_id=1000 group_id=100]
450        21          0       49      /      /sys/fs/fuse/connections          rw,relatime                       [shared:357]      fusectl           fusectl                   [rw]
496        45          0       11      /      /sys/kernel/debug/tracing         rw,relatime                       [shared:400]      tracefs           tracefs                   [rw]
241        382         0       64      /      /run/user/1000/doc                rw,nosuid,nodev,relatime          [shared:141]      fuse              /dev/fuse                 [rw user_id=1000 group_id=100]
```
