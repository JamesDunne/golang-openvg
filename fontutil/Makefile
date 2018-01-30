all: font2openvg

clean:
	rm -f font2openvg font2openvg.o

CPP     = g++
CFLAGS  += -c -Wall $(shell freetype-config --cflags)
LDFLAGS += -Wall $(shell freetype-config --libs)

font2openvg.o: font2openvg.cpp
	$(CPP) -o $@ $(CFLAGS) $<

font2openvg: font2openvg.o
	$(CPP) -o $@ $(LDFLAGS) $<
