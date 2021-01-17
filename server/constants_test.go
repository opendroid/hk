package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart_Constants(t *testing.T) {
	assert.Equal(t, "public/css", css)
	assert.Equal(t, "public/images", images)
	assert.Equal(t, "public/templates", templates)
	assert.Equal(t, "/public/css", cssAbs)
	assert.Equal(t, "/public/images", imagesAbs)
	assert.Equal(t, "/public/templates", templatesAbs)
	assert.Equal(t, "public/templates/*.gohtml", htmlGlob)
	assert.Equal(t, "public/robots.txt", robots)
	assert.Equal(t, "public/sitemap.xml", sitemap)
	assert.Equal(t, "public/images/icons/favicon-16x16.png", favicon)
	assert.Equal(t, "public/images/icons/favicon-32x32.png", favicon32)
}
