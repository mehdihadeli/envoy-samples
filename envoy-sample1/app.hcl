container "api1" {
    image {
        name = "nicholasjackson/fake-service:v0.8.0"
    }
     port {
	    local = "9090"
		host = "9091"
		remote = "9091"
	 }
    env {
        key = "NAME"
        value = "API 1"
    }
}

container "api2" {
    image {
        name = "nicholasjackson/fake-service:v0.8.0"
    }
     port {
	    local = "9090"
		host = "9092"
		remote = "9092"
	 }
    env {
        key = "NAME"
        value = "API 2"
    }
}



container "api3" {
    image {
        name = "nicholasjackson/fake-service:v0.8.0"
    }
     port {
	    local = "9090"
		host = "9093"
		remote = "9093"
	 }
    env {
        key = "NAME"
        value = "API 3"
    }
}


container "api4" {
    image {
        name = "nicholasjackson/fake-service:v0.8.0"
    }
     port {
	    local = "9090"
		host = "9094"
		remote = "9094"
	 }
    env {
        key = "NAME"
        value = "API 4"
    }
}

