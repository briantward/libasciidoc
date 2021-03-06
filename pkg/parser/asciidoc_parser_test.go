package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func verify(t GinkgoTInterface, expectedResult interface{}, content string, options ...parser.Option) {
	log.Debugf("processing: %s", content)
	reader := strings.NewReader(content)
	allOptions := append(options)
	result, err := parser.ParseReader("", reader, allOptions...)
	if err != nil {
		log.WithError(err).Errorf("Error found while parsing the document (%T)", err)
	}
	require.NoError(t, err)
	t.Logf("actual document: `%s`", spew.Sdump(result))
	t.Logf("expected document: `%s`", spew.Sdump(expectedResult))
	assert.EqualValues(t, expectedResult, result)
}

func verifyError(t GinkgoTInterface, content string, options ...parser.Option) {
	log.Debugf("processing: %s", content)
	reader := strings.NewReader(content)
	allOptions := append(options, parser.Recover(false))
	_, err := parser.ParseReader("", reader, allOptions...)
	require.Error(t, err)
}
