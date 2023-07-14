package media

import (
	"time"
	"unsafe"

	"github.com/thinkonmay/thinkshare-daemon/persistent/gRPC/packet"
)

/*
#include <gst/gst.h>
#include <stdio.h>



typedef struct _Soundcard {
    char device_id[500];
    char name[500];
    char api[50];

    gboolean isdefault;
    gboolean loopback;

    int active;
}Soundcard;

typedef struct _Mic {
    char device_id[500];
    char name[500];
    char api[50];

    gboolean isdefault;
    gboolean loopback;

    int active;
}Mic;

typedef struct _Monitor {
    guint64 monitor_handle;
    int primary;

    char device_name[100];
    char adapter[100];
    char monitor_name[100];

    int width, height;

    int active;
}Monitor;

typedef struct _MediaDevice
{
    Monitor monitors[10];
    Soundcard soundcards[10];
    Mic microphones[10];
}MediaDevice;

static void
device_foreach(GstDevice* device,
               gpointer data)
{
    MediaDevice* source = (MediaDevice*) data;
    gchar* klass = gst_device_get_device_class(device);

    // handle monitor
    if(!g_strcmp0(klass,"Source/Monitor")) {
        GstStructure* device_structure = gst_device_get_properties(device);
        gchar* api = (gchar*)gst_structure_get_string(device_structure,"device.api");
        if(g_strcmp0(api,"d3d11")) {
            g_object_unref(device);
            return;
        }

        int i = 0;
        while (source->monitors[i].active) { i++; }
        Monitor* monitor = &source->monitors[i];

        gchar* name = gst_device_get_display_name(device);
        memcpy(monitor->monitor_name,name,strlen(name));
        monitor->active = TRUE;

        gchar* adapter = (gchar*)gst_structure_get_string(device_structure,"device.adapter.description");
        memcpy(monitor->adapter,adapter,strlen(adapter));

        gchar* device_name = (gchar*)gst_structure_get_string(device_structure,"device.name");
        memcpy(monitor->device_name,device_name,strlen(device_name));

        int top, left, right, bottom = 0;
        gst_structure_get_int(device_structure,"display.coordinates.right",&right);
        gst_structure_get_int(device_structure,"display.coordinates.top",&top);
        gst_structure_get_int(device_structure,"display.coordinates.left",&left);
        gst_structure_get_int(device_structure,"display.coordinates.bottom",&bottom);

        monitor->width =  right - left;
        monitor->height = bottom - top;


        gst_structure_get_uint64(device_structure,"device.hmonitor",&monitor->monitor_handle);
        gst_structure_get_boolean(device_structure,"device.primary",&monitor->primary);
    }

    // handle audio
    if(!g_strcmp0(klass,"Audio/Source")) {
        GstStructure* device_structure = gst_device_get_properties(device);
        gchar* api = (gchar*)gst_structure_get_string(device_structure,"device.api");
        if(!g_strcmp0(api,"wasapi")) {
            int i = 0;
            while (source->soundcards[i].active) { i++; }
            Soundcard* soundcard = &source->soundcards[i];

            memcpy(soundcard->api,api,strlen(api));

            gchar* name = gst_device_get_display_name(device);
            memcpy(soundcard->name,name,strlen(name));
            soundcard->active = TRUE;

            gchar* device_name = (gchar*)gst_structure_get_string(device_structure,"wasapi.device.description");
            memcpy(soundcard->name,device_name,strlen(device_name));

            gchar* strid = (gchar*)gst_structure_get_string(device_structure,"device.strid");
            memcpy(soundcard->device_id,strid,strlen(strid));
        } else if (!g_strcmp0(api,"wasapi2")) {
            int i = 0;
            while (source->soundcards[i].active) { i++; }
            Soundcard* soundcard = &source->soundcards[i];

            memcpy(soundcard->api,api,strlen(api));
            gst_structure_get_boolean(device_structure,"device.default",&soundcard->isdefault);
            gst_structure_get_boolean(device_structure,"wasapi2.device.loopback",&soundcard->loopback);

            gchar* name = gst_device_get_display_name(device);
            memcpy(soundcard->name,name,strlen(name));
            soundcard->active = TRUE;


            gchar* strid = (gchar*)gst_structure_get_string(device_structure,"device.id");
            memcpy(soundcard->device_id,strid,strlen(strid));
        } else {
            g_object_unref(device);
            return;
        }
    }

    // handle audio
    if(!g_strcmp0(klass,"Audio/Sink")) {
        GstStructure* device_structure = gst_device_get_properties(device);
        gchar* api = (gchar*)gst_structure_get_string(device_structure,"device.api");
        if(!g_strcmp0(api,"wasapi")) {
            int i = 0;
            while (source->microphones[i].active) { i++; }
            Mic* mic = &source->microphones[i];

            gchar* strid = (gchar*)gst_structure_get_string(device_structure,"device.strid");
            memcpy(mic->device_id,strid,strlen(strid));
            memcpy(mic->api,api,strlen(api));
            gchar* name = gst_device_get_display_name(device);
            memcpy(mic->name,name,strlen(name));

            mic->active = TRUE;
        } else {
            g_object_unref(device);
            return;
        }
    }
    g_object_unref(device);
}



MediaDevice*
query_media_device()
{
    static MediaDevice dev;
    memset(&dev,0,sizeof(MediaDevice));

    gst_init(NULL, NULL);
    GstDeviceMonitor* monitor = gst_device_monitor_new();
    if(!gst_device_monitor_start(monitor)) {
        return (void*)"fail to start device monitor";
    }

    GList* device_list = gst_device_monitor_get_devices(monitor);
    g_list_foreach(device_list,(GFunc)device_foreach,&dev);

    return &dev;
}

#cgo pkg-config: gstreamer-1.0
*/
import "C"

type DeviceQuery unsafe.Pointer

func GetDevice() *packet.MediaDevice {
	result := &packet.MediaDevice{
		Timestamp: time.Now().Format(time.RFC3339),
	}
	query := C.query_media_device()

    for _,mic := range query.microphones {
        if mic.active == 0 {
            continue
        }
        result.Microphones = append(result.Microphones, &packet.Microphone{
			Name:       C.GoString(&mic.name[0]),
			DeviceID:   C.GoString(&mic.device_id[0]),
			Api:        C.GoString(&mic.api[0]),
        })
    }

    for _,sound := range query.soundcards {
        if sound.active == 0 {
            continue
        }
        result.Soundcards = append(result.Soundcards, &packet.Soundcard{
			Name:       C.GoString(&sound.name[0]),
			DeviceID:   C.GoString(&sound.device_id[0]),
			Api:        C.GoString(&sound.api[0]),
        })
    }

    for _,monitor := range query.monitors {
        if monitor.active == 0 {
            continue
        }
        result.Monitors = append(result.Monitors, &packet.Monitor{
			Framerate:     60,
			MonitorHandle: int32(monitor.monitor_handle),
			Width:         int32(monitor.width),
			Height:        int32(monitor.height),
			MonitorName:   C.GoString(&monitor.monitor_name[0]),
			Adapter:       C.GoString(&monitor.adapter[0]),
			DeviceName:    C.GoString(&monitor.device_name[0]),
			IsPrimary:     monitor.primary == 1,
        })
    }


	return result
}

