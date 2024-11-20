# memo

2024/11/12現在、[ここ](https://github.com/riscv-software-src/riscv-pk/pull/114)を参考にして`stub-lp64.h`をコピーしておかないとpkのコンパイルができない
あとは[ここ](https://zenn.dev/ohno418/articles/5f6d5e01dc4981) を参考にする

## 実行方法

```bash
riscv64-unknown-linux-gnu-gcc -static fib.c -o fib
spike $RISCV/riscv64-unknown-linux-gnu/bin/pk fib
```

## qemu-riscv64で実行する方法

### 事前にqemu-userをインストールする

```bash
sudo apt install qemu-user
```

```bash
docker exec -ti spike  bash -c "/opt/riscv/bin/riscv64-unknown-linux-gnu-gcc -static ~/fib.c && mv a.out /mnt/"
qemu-riscv64 a.out
```

