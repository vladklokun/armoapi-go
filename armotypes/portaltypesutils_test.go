package armotypes

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/francoispqt/gojay"
	"github.com/stretchr/testify/assert"
)

func TestAttributesDesignatorsFromWLID(t *testing.T) {
	attDesig := AttributesDesignatorsFromWLID("wlid://cluster-liortest1/namespace-default/deployment-payment")
	if attDesig.Attributes[AttributeCluster] != "liortest1" ||
		attDesig.Attributes[AttributeNamespace] != "default" ||
		attDesig.Attributes[AttributeKind] != "deployment" ||
		attDesig.Attributes[AttributeName] != "payment" {
		t.Errorf("wrong attributes desigantors:%v", attDesig)
	}

	attDesig = AttributesDesignatorsFromWLID("wlid://cluster-liortest1/namespace-default/deployment")
	if attDesig.Attributes[AttributeCluster] != "liortest1" ||
		attDesig.Attributes[AttributeNamespace] != "default" ||
		attDesig.Attributes[AttributeKind] != "deployment" {
		t.Errorf("wrong attributes desigantors:%v", attDesig)
	}
	attDesig = AttributesDesignatorsFromWLID("wlid://cluster-liortest1/namespace-default/")
	if attDesig.Attributes[AttributeCluster] != "liortest1" ||
		attDesig.Attributes[AttributeNamespace] != "default" {
		t.Errorf("wrong attributes desigantors:%v", attDesig)
	}
	attDesig = AttributesDesignatorsFromWLID("wlid://cluster-liortest1")
	if attDesig.Attributes[AttributeCluster] != "liortest1" {
		t.Errorf("wrong attributes desigantors:%v", attDesig)
	}
}

//go:embed fixtures/designatorTestCase.json
var designatorTestCase string

func TestDesignatorDecoding(t *testing.T) {
	designator := &PortalDesignator{}
	er := gojay.NewDecoder(strings.NewReader(designatorTestCase)).DecodeObject(designator)
	if er != nil {
		t.Errorf("decode failed due to: %v", er.Error())
	}
	assert.Equal(t, DesignatorAttributes, designator.DesignatorType)
	assert.Equal(t, "myCluster", designator.Attributes[AttributeCluster])
	assert.Equal(t, "8190928904639901517", designator.Attributes[AttributeWorkloadHash])
	assert.Equal(t, "myName", designator.Attributes[AttributeName])
	assert.Equal(t, "myNS", designator.Attributes[AttributeNamespace])
	assert.Equal(t, "deployment", designator.Attributes[AttributeKind])
	assert.Equal(t, "e57ec5a0-695f-4777-8366-1c64fada00a0", designator.Attributes[AttributeCustomerGUID])
	assert.Equal(t, "myContainer", designator.Attributes[AttributeContainerName])
}

func TestAttributesDesignatorsFromImageTag(t *testing.T) {
	deisgs := AttributesDesignatorsFromImageTag("docker.elastic.co/elasticsearch/elasticsearch:7.9.2")

	assert.Equal(t, "docker.elastic.co/elasticsearch", deisgs.Attributes[AttributeRegistryName])
	assert.Equal(t, "elasticsearch", deisgs.Attributes[AttributeRepository])
	assert.Equal(t, "7.9.2", deisgs.Attributes[AttributeTag])
	assert.Equal(t, 3, len(deisgs.Attributes))

	deisgs = AttributesDesignatorsFromImageTag("docker.elastic.co/elasticsearch/elasticsearch")

	assert.Equal(t, "docker.elastic.co/elasticsearch", deisgs.Attributes[AttributeRegistryName])
	assert.Equal(t, "elasticsearch", deisgs.Attributes[AttributeRepository])
	assert.Equal(t, "", deisgs.Attributes[AttributeTag])
	assert.Equal(t, 2, len(deisgs.Attributes))

	deisgs = AttributesDesignatorsFromImageTag("docker.elastic.co/elasticsearch")

	assert.Equal(t, "docker.elastic.co", deisgs.Attributes[AttributeRegistryName])
	assert.Equal(t, "elasticsearch", deisgs.Attributes[AttributeRepository])
	assert.Equal(t, "", deisgs.Attributes[AttributeTag])
	assert.Equal(t, 2, len(deisgs.Attributes))

	deisgs = AttributesDesignatorsFromImageTag("docker.elastic.co/")

	assert.Equal(t, "docker.elastic.co", deisgs.Attributes[AttributeRegistryName])
	assert.Equal(t, "", deisgs.Attributes[AttributeRepository])
	assert.Equal(t, "", deisgs.Attributes[AttributeTag])
	assert.Equal(t, 1, len(deisgs.Attributes))

	deisgs = AttributesDesignatorsFromImageTag("docker.elastic.co")

	assert.Equal(t, "docker.elastic.co", deisgs.Attributes[AttributeRegistryName])
	assert.Equal(t, "", deisgs.Attributes[AttributeRepository])
	assert.Equal(t, "", deisgs.Attributes[AttributeTag])
	assert.Equal(t, 1, len(deisgs.Attributes))

	deisgs = AttributesDesignatorsFromImageTag("")

	assert.Equal(t, "", deisgs.Attributes[AttributeRegistryName])
	assert.Equal(t, "", deisgs.Attributes[AttributeRepository])
	assert.Equal(t, "", deisgs.Attributes[AttributeTag])
	assert.Equal(t, 1, len(deisgs.Attributes))
}
