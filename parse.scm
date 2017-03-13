; s = struct
; f = field
(define (parse-meta s f)
  (cond
    ((equal? f 'name)         (car s))
    ((equal? f 'version)      (cadr s))
    ((equal? f 'dependencies) (caddr s))
    ((equal? f 'replaces)     (cadddr s))
    (else (Nothing))))

; s = struct
; f = field
(define (parse-source s f)
  (cond
    ((equal? f 'download)    (car s))
    ((equal? f 'local-files) (cadr s))
    ((equal? f 'sha256)      (caddr s))
    (else (Nothing))))

; s = struct
; f = field
(define (parse-feature s f)
  (cond
    ((equal? f 'name)         (car s))
    ((equal? f 'default?)     (cadr s))
    ((equal? f 'enable-flag)  (caddr s))
    ((equal? f 'disable-flag) (cadddr s))
    (else (Nothing))))

; s = struct
; n = nth feature to get
(define (get-nth-feature s n)
  (let aux ((in s) (nth n) (out '()))
    (cond
      ((null? in)  (Nothing))
      ((zero? nth) (car in))
      (else (aux (cdr in) (- nth 1) out)))))

; s = struct
; f = field
(define (parse-extra s f)
  (cond
    ((equal? f 'configure-stage) (car s))
    ((equal? f 'pre-build)       (cadr s))
    ((equal? f 'post-build)      (caddr s))
    (else (Nothing))))
