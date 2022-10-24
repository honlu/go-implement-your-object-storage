module apiServer

go 1.13

replace lib/rabbitmq => ../../src/lib/rabbitmq

replace lib/objectstream => ../../src/lib/objectstream

replace lib/es => ../../src/lib/es

replace lib/utils => ../../src/lib/utils

require (
	lib/es v0.0.0-00010101000000-000000000000
	lib/objectstream v0.0.0-00010101000000-000000000000
	lib/rabbitmq v0.0.0-00010101000000-000000000000
	lib/utils v0.0.0-00010101000000-000000000000
)
