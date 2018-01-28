#include <bcm_host.h>
#include <EGL/egl.h>
#include <stdint.h>

EGLint
create(uint32_t* width, uint32_t* height, EGLDisplay* display, EGLSurface* surface)
{
	int32_t success = 0;
	EGLBoolean result;
	EGLint num_config;

	static EGL_DISPMANX_WINDOW_T nativewindow;

	DISPMANX_ELEMENT_HANDLE_T dispman_element;
	DISPMANX_DISPLAY_HANDLE_T dispman_display;
	DISPMANX_UPDATE_HANDLE_T dispman_update;
	VC_RECT_T dst_rect;
	VC_RECT_T src_rect;
	static VC_DISPMANX_ALPHA_T alpha = {
		DISPMANX_FLAGS_ALPHA_FIXED_ALL_PIXELS,
		255, 0
	};

	static const EGLint attribute_list[] = {
		EGL_RED_SIZE, 8,
		EGL_GREEN_SIZE, 8,
		EGL_BLUE_SIZE, 8,
		EGL_ALPHA_SIZE, 8,
		EGL_SURFACE_TYPE, EGL_WINDOW_BIT,
		EGL_NONE
	};

	EGLConfig config;
	EGLContext context;

	bcm_host_init();

	// get an EGL display connection
	*display = eglGetDisplay(EGL_DEFAULT_DISPLAY);
	if (*display == EGL_NO_DISPLAY) return eglGetError();

	// initialize the EGL display connection
	result = eglInitialize(*display, NULL, NULL);
	if (EGL_FALSE == result) return eglGetError();

	// bind OpenVG API
	result = eglBindAPI(EGL_OPENVG_API);
	if (EGL_FALSE == result) return eglGetError();

	// get an appropriate EGL frame buffer configuration
	result = eglChooseConfig(*display, attribute_list, &config, 1, &num_config);
	if (EGL_FALSE == result) return eglGetError();

	// create an EGL rendering context
	context = eglCreateContext(*display, config, EGL_NO_CONTEXT, NULL);
	if (EGL_NO_CONTEXT == context) return eglGetError();

	// create an EGL window surface
	success = graphics_get_display_size(0 /* LCD */, width, height);
	if (success < 0) return -1;

	vc_dispmanx_rect_set(&dst_rect, 0, 0, *width, *height);
	vc_dispmanx_rect_set(&src_rect, 0 << 16, 0 << 16, *width << 16, *height << 16);

	dispman_display = vc_dispmanx_display_open(0 /* LCD */ );
	dispman_update = vc_dispmanx_update_start(0);

	dispman_element = vc_dispmanx_element_add(dispman_update, dispman_display, 0 /*layer */ , &dst_rect, 0 /*src */ ,
						  &src_rect, DISPMANX_PROTECTION_NONE, &alpha, 0 /*clamp */ ,
						  0 /*transform */ );

	nativewindow.element = dispman_element;
	nativewindow.width = *width;
	nativewindow.height = *height;
	vc_dispmanx_update_submit_sync(dispman_update);

	*surface = eglCreateWindowSurface(*display, config, &nativewindow, NULL);
	if (EGL_NO_SURFACE == *surface) return eglGetError();

	// preserve the buffers on swap
	result = eglSurfaceAttrib(*display, *surface, EGL_SWAP_BEHAVIOR, EGL_BUFFER_PRESERVED);
	if (EGL_FALSE == result) return eglGetError();

	// connect the context to the surface
	result = eglMakeCurrent(*display, *surface, *surface, context);
	if (EGL_FALSE == result) return eglGetError();

	return EGL_SUCCESS;
}
