# Maintainer: Benjamin Lopez <contact@scotow.com>

pkgname=notigo
pkgver=1.0.0
pkgrel=1
pkgdesc="Send iOS/Android notifications using IFTTT's Webhook"
arch=('x86_64')
url="https://github.com/scotow/${pkgname}"
license=('MIT')
makedepends=('go' 'git')
source=("${pkgname}-${pkgver}.tar.gz::https://github.com/scotow/${pkgname}/archive/${pkgver}.tar.gz")
sha256sums=('28c59c2e91f5428debfda41b5a76317db70fe36326c49cf45480cde84d5637e0')

prepare(){
  mkdir -p src/github.com/scotow
  ln -rTsf "${pkgname}-${pkgver}" "src/github.com/scotow/${pkgname}"
}

build(){
  export GOPATH="${srcdir}"
  cd "src/github.com/scotow/${pkgname}"
  go install \
	-gcflags "all=-trimpath=${GOPATH}/src" \
	-asmflags "all=-trimpath=${GOPATH}/src" \
	-ldflags "-extldflags ${LDFLAGS}" \
	./cmd/...
}

package(){
  install -Dm755 "bin/${pkgname}" "${pkgdir}/usr/bin/${pkgname}"

  cd "${pkgname}-${pkgver}"
  install -Dm644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
