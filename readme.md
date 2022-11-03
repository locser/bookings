# Bookings and Reservations

The repository for [Building Modern Web Applications with Go](https://www.udemy.com/course/building-modern-web-applications-with-go/?referralCode=0415FB906223F10C6800).

- Built in Go version 1.19
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)

- Uses [sweetalert2](sweetalert2.github.io/#examples)
- Uses [Govalidator](https://github.com/asaskevich/govalidator)

How To Run:

- Test main : cd /cmd/web  -> go test -v
- Test Handler: cd /bookings/handler/ -> go test -v 

- Run:  open Terminal: go run ./cmd/web, or 
+ ./run.sh (for MAC)
+ .\run.bat(WINDOW)