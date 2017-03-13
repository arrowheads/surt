(define meta
  '((name "foobar")
    (version "13.3.7")
    (dependencies "x" "y" "z")
    (replaces "barfoo")))

(define source
  '((download
       "https://anexcellenturl.com/anexcellentprogram.tar.gz"
       "https://anexcellenturl.com/anexcellentpatch.gz")
    (local-files
      "patch1.patch"
      "patch2.patch"
      "patch3.patch")
    (sha256
      "fooooooooooooooooooooooooooooooooooooooooooooooooooo"
      "baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaar"
      "baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaz")))

(define options
  '(strip-binaries
    static-binaries
    parallel-build))

(define features
  '((foo (default #t)
          (enable-flag "--enable-foo")
          (disable-flag "--disable-foo")
          (requires (Nothing)))
    (bar (default #t)
         (enable-flag "--enable-bar")
         (disable-flag "--disable-bar")
         (requires foo))))

(define extra
  '((configure-stage
       "--prefix=/usr"
       "--execprefix=/usr")
    (pre-build
      "mkdir foo")
    (post-build
      "rm bar")))
