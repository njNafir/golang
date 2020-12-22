module hello

go 1.15

replace example.com/greetings => ../module/greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
