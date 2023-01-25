# Exercise 1 - Bootable Linux image via QEMU

## Description
Builds a bootable `x86_64` Linux 6.1 image, with a personalized file
(`hello_world.c`) as PID1.

`hello_world.c` prints to output `hello world` and then hangs.

## Dependencies
`wget cpio gcc tar qemu`

## Usage
Run `make run` on the root folder. 

You can adjust some settings by changing the variables at the top of
the `Makefile`.

NOTE: I parallelized the building of the Linux kernel. By default it uses 8 cores.
If you want to change the number of cores used, please edit the `N_CORES` variable on
the beginning of the `Makefile`.


