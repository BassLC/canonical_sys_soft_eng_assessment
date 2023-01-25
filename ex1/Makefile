PID1_CODE_FILE=hello_world.c
N_CORES=8
LINUX_FOLDER=linux-6.1/
KERNEL_IMG=$(LINUX_FOLDER)/arch/x86_64/boot/bzImage

.PHONY: run clean

run: initramfs $(KERNEL_IMG)
	@echo "> Running QEMU"
	qemu-system-x86_64 -serial stdio \
                   -kernel linux-6.1/arch/x86_64/boot/bzImage \
                   -initrd initramfs \
                   -m 128
$(LINUX_FOLDER):
	@echo "> Getting the 6.1 kernel"
	wget -nv --show-progress https://cdn.kernel.org/pub/linux/kernel/v6.x/linux-6.1.tar.gz
	tar xf linux-6.1.tar.gz


$(KERNEL_IMG): $(LINUX_FOLDER)
	@echo "> Building the kernel image"
	$(MAKE) -C $(LINUX_FOLDER) defconfig
	$(MAKE) -C $(LINUX_FOLDER) -j $(N_CORES)

initramfs: $(PID1_CODE_FILE)
	@echo "> Building initramfs"
	gcc -static $(PID1_CODE_FILE) -o init
	echo init | cpio -o --format=newc -O initramfs

clean:
	rm -r $(LINUX_FOLDER) linux-6.1.tar.gz initramfs init
