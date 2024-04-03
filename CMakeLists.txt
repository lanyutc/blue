cmake_minimum_required(VERSION 3.25)

project(blue)

function(build_blue name srcdir)
	message(STATUS "build: ${name} ${srcdir}")
	add_custom_target(${name})
	add_custom_command(TARGET ${name}
		COMMAND go build -o "${CMAKE_CURRENT_BINARY_DIR}/${name}" 
		WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR}
	)

	add_custom_target(${name}_exe ALL DEPENDS ${name})
endfunction(build_blue)

add_subdirectory(naming)
add_subdirectory(pid_dispatch)
add_subdirectory(example/netexample/client)
add_subdirectory(example/netexample/server)
add_subdirectory(example/rpcexample/client)
add_subdirectory(example/rpcexample/server)
add_subdirectory(example/uidexample)
add_subdirectory(example/netexample/udp_server)
add_subdirectory(example/netexample/udp_client)
