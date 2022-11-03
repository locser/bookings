go build -o bookings.exe ./cmd/web/. || exit /b
bookings.exe

::Had to use the || exit /b otherwise it would just run the previously created .exe file