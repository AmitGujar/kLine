#!/bin/bash

# Change into the directory where the .deb files are located
# cp /home/ubuntu/*.deb .
cd dists/stable/main/binary-amd64
# rm *.deb
# cp /home/ubuntu/*.deb .
# Generate Packages files
apt-ftparchive packages . >Packages
bzip2 -c9 Packages >Packages.bz2
xz -c9 Packages >Packages.xz
xz -5fkev --format=lzma Packages >Packages.lzma
lz4 -c9 Packages >Packages.lz4
gzip -c9 Packages >Packages.gz
zstd -c19 Packages >Packages.zst

# Generate Contents files
apt-ftparchive contents . >Contents-kLine-amd64
bzip2 -c9 Contents-kLine-amd64 >Contents-kLine-amd64.bz2
xz -c9 Contents-kLine-amd64 >Contents-kLine-amd64.xz
xz -5fkev --format=lzma Contents-kLine-amd64 >Contents-kLine-amd64.lzma
lz4 -c9 Contents-kLine-amd64 >Contents-kLine-amd64.lz4
gzip -c9 Contents-kLine-amd64 >Contents-kLine-amd64.gz
zstd -c19 Contents-kLine-amd64 >Contents-kLine-amd64.zst

# Change back to the root of the repository
# cd /var/www/html/kline/dists/stable
# cd ../..
cd ../..
# Generate Release files
grep -E "Origin:|Label:|Suite:|Version:|Codename:|Architectures:|Components:|Description:" Release >Base
# grep -E "Origin:|Label:|Suite:|Version:|Codename:|Architectures:|Components:|Description:" Release >../../Base
apt-ftparchive release . >Release
#gpg --clearsign -o InRelease Release
cat Base Release >out && mv out Release

# Sign the Release file
gpg -abs -u 9D0FA86CF53F15831B6D8CB69633467FEBF46E4E -o Release.gpg Release
