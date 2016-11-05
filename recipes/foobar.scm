(meta
  (name "foobar")
  (version "13.3.7")
  (dependencies ("x" "y" "z"))
  (replaces '("barfoo")))

(source
  (download
    "https://anexcellenturl.com/anexcellentprogram.tar.gz")
  (local-files
    "patch1.patch"
    "patch2.patch"
    "patch3.patch")
  (sha256
    "fooooooooooooooooooooooooooooooooooooooooooooooooooo"
    "baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaar"
    "baaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaz"))

(options
  'strip-binaries
  'static-binaries
  'parallel-build)

(features
  (define foo
    ('(default true)
     '(enable-flag "--enable-foo")
     '(disable-flag "--disable-foo")
     (requires
       '(bar))))
  (define bar
    ('(default true)
     '(enable-flag "--enable-bar")
     '(disable-flag "--disable-bar")
     (requires
       '()))))

(extra
  (configure-stage
     "--prefix=/usr"
     "--execprefix=/usr")
  (pre-build
     "mkdir foo")
  (post-build
    "rm bar"))

