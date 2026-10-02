---
page_title: "infra"
subcategory: ""
description: "InfraMetadata stores information about instance infrastructure."
xcsh_docs: {"aliases": ["infra"], "body_bytes": 11159, "body_sha256": "sha256:dd4b028eb6857cd938447862a4bceeedb1cd5a9cab7b59ea21d39e04be6845ef", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:registration:properties:infra:bond_config", "xcsh-docs:resources:registration:properties:infra:hugepages", "xcsh-docs:resources:registration:properties:infra:hw_info", "xcsh-docs:resources:registration:properties:infra:interfaces", "xcsh-docs:resources:registration:properties:infra:internet_proxy", "xcsh-docs:resources:registration:properties:infra:sw_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "id": "xcsh-docs:resources:registration:properties:infra", "parent_id": "xcsh-docs:resources:registration:reference", "path": "documentation/resources/registration/properties/infra/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1110013003021213-1032113012313300-0113333001130023-3322230323211220-0122200322221221-3020222111323313-3002002332323333-3332103312113111", "registry_path": "docs/guides/resources--registration--reference--group-001.md", "relationships": [{"anchor": "schema-infra--hostname", "enforcement": "provider-schema", "group": "infra:RequiredObjectAttributes:hostname,interfaces", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:infra", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "infra:RequiredObjectAttributes:hostname,interfaces", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:infra:interfaces", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["infra"], "schema_version": 1, "sections": [{"aliases": ["availability zone"], "anchor": "schema-infra--availability_zone", "description": "An Availability Zone is a high-availability offering that protects your applications and data from datacenter failures.", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "availability_zone"], "syntax": "attribute", "type": "string"}, {"aliases": ["bond config"], "anchor": "section", "description": "Bond device configuration for VPM registration.", "document_id": "xcsh-docs:resources:registration:properties:infra:bond_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-infra--bond_config--interfaces", "enforcement": "provider-schema", "group": "infra.bond_config:RequiredObjectAttributes:interfaces,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:infra:bond_config", "type": "requires"}, {"anchor": "schema-infra--bond_config--name", "enforcement": "provider-schema", "group": "infra.bond_config:RequiredObjectAttributes:interfaces,name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:registration:properties:infra:bond_config", "type": "requires"}], "schema_path": ["infra", "bond_config"], "syntax": "block", "type": "object"}, {"aliases": ["certified hw"], "anchor": "schema-infra--certified_hw", "description": "Certified HW name used to map with F5XC certified_hardware definition.", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["domain"], "anchor": "schema-infra--domain", "description": "Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5 Distributed Cloud.", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["hostname"], "anchor": "schema-infra--hostname", "description": "Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["hugepages"], "anchor": "section", "description": "Hugepage settings for CE on K8s SMV2 site.", "document_id": "xcsh-docs:resources:registration:properties:infra:hugepages", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hugepages"], "syntax": "block", "type": "object"}, {"aliases": ["hw info"], "anchor": "section", "description": "OsInfo holds information about host OS and HW.", "document_id": "xcsh-docs:resources:registration:properties:infra:hw_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info"], "syntax": "block", "type": "object"}, {"aliases": ["instance id"], "anchor": "schema-infra--instance_id", "description": "Instance ID (assigned by infrastructure provider)", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "instance_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["interfaces"], "anchor": "section", "description": "Machine interfaces present during registration time.", "document_id": "xcsh-docs:resources:registration:properties:infra:interfaces", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "interfaces"], "syntax": "block", "type": "object"}, {"aliases": ["internet proxy"], "anchor": "section", "description": "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.", "document_id": "xcsh-docs:resources:registration:properties:infra:internet_proxy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "internet_proxy"], "syntax": "block", "type": "object"}, {"aliases": ["is slo static"], "anchor": "schema-infra--is_slo_static", "description": "Indicates whether the SLO is static.", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "is_slo_static"], "syntax": "attribute", "type": "bool"}, {"aliases": ["machine id"], "anchor": "schema-infra--machine_id", "description": "Machine ID - generated by operating system.", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "machine_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["provider ref"], "anchor": "schema-infra--provider_ref", "description": "Infrastructure provider enum for registration. It describes where is instance running. Provider was not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other provider, which was not identified by system. Possible values are `UNKNOWN`, `AWS`, `GOOGLE`, `AZURE`, `VMWARE`, `KVM`,", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["sw info"], "anchor": "section", "description": "SWInfo holds information about sw version.", "document_id": "xcsh-docs:resources:registration:properties:infra:sw_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "sw_info"], "syntax": "block", "type": "object"}, {"aliases": ["timestamp"], "anchor": "schema-infra--timestamp", "description": "It's used to verify machine have acceptable time difference from server.", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["zone"], "anchor": "schema-infra--zone", "description": "Instance zone (or region), depends on provider.", "document_id": "xcsh-docs:resources:registration:properties:infra", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "zone"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/properties/infra/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "InfraMetadata stores information about instance infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- infra

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

InfraMetadata stores information about instance infrastructure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hostname",
    "interfaces")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
infra {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-infra--availability_zone"></a>

### availability_zone property

Type: `"string"`. Optional.

Availability Zone is a high-availability offering that protects your applications and data from
datacenter failures.

Upstream description:

An Availability Zone is a high-availability offering that protects your applications and data from
datacenter failures.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/bond_config/): complete subsection reference.

<a id="schema-infra--certified_hw"></a>

### certified_hw property

Type: `"string"`. Optional.

Certified HW name used to map with F5XC certified\_hardware definition.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--domain"></a>

### domain property

Type: `"string"`. Optional.

Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5
Distributed Cloud.

Upstream description:

Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5
Distributed Cloud.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--hostname"></a>

### hostname property

Type: `"string"`. Optional.

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Upstream description:

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hugepages/): complete subsection reference.

- [hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/): complete subsection reference.

<a id="schema-infra--instance_id"></a>

### instance_id property

Type: `"string"`. Optional.

Instance ID (assigned by infrastructure provider).

Upstream description:

Instance ID (assigned by infrastructure provider)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/interfaces/): complete subsection reference.

- [internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/internet_proxy/): complete subsection reference.

<a id="schema-infra--is_slo_static"></a>

### is_slo_static property

Type: `"bool"`. Optional.

Is SLO Static. Indicates whether the SLO is static.

Upstream description:

Indicates whether the SLO is static.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--machine_id"></a>

### machine_id property

Type: `"string"`. Optional.

Machine ID - generated by operating system.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--provider_ref"></a>

### provider_ref property

Type: `"string"`. Optional.

\[Enum:
UNKNOWN|AWS|GOOGLE|AZURE|VMWARE|KVM|OTHER|VOLTERRA|IBMCLOUD|UNKNOWN\_K8S|AWS\_K8S|GCP\_K8S|AZURE\_K8S|VMWARE\_K8S|KVM\_K8S|OTHER\_K8S|VOLTERRA\_K8S|IBMCLOUD\_K8S|F5OS|RSERIES|OCI|NUTANIX|OPENSTACK|EQUINIX|OPENSHIFT\_VIRTUALIZATION|KUBERNETES\]
Infrastructure provider enum for registration. It describes where is instance running. Provider was
not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other
provider, which was not identified by system. Possible values are \`UNKNOWN\`, \`AWS\`, \`GOOGLE\`,
\`AZURE\`, \`VMWARE\`, \`KVM\`, \`OTHER\`, \`VOLTERRA\`, \`IBMCLOUD\`, \`UNKNOWN\_K8S\`,
\`AWS\_K8S\`, \`GCP\_K8S\`, \`AZURE\_K8S\`, \`VMWARE\_K8S\`, \`KVM\_K8S\`, \`OTHER\_K8S\`,
\`VOLTERRA\_K8S\`, \`IBMCLOUD\_K8S\`, \`F5OS\`, \`RSERIES\`, \`OCI\`, \`NUTANIX\`, \`OPENSTACK\`,
\`EQUINIX\`, \`OPENSHIFT\_VIRTUALIZATION\`, \`KUBERNETES\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNKNOWN",
    "AWS",
    "GOOGLE",
    "AZURE",
    "VMWARE",
    "KVM",
    "OTHER",
    "VOLTERRA",
    "IBMCLOUD",
    "UNKNOWN_K8S",
    "AWS_K8S",
    "GCP_K8S",
    "AZURE_K8S",
    "VMWARE_K8S",
    "KVM_K8S",
    "OTHER_K8S",
    "VOLTERRA_K8S",
    "IBMCLOUD_K8S",
    "F5OS",
    "RSERIES",
    "OCI",
    "NUTANIX",
    "OPENSTACK",
    "EQUINIX",
    "OPENSHIFT_VIRTUALIZATION",
    "KUBERNETES"),
}
```

- [sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/sw_info/): complete subsection reference.

<a id="schema-infra--timestamp"></a>

### timestamp property

Type: `"string"`. Optional.

It's used to verify machine have acceptable time difference from server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "temporal",
    "constraintType": "string",
    "deterministic": true,
    "format": "date-time",
    "formatDescription": "ISO 8601 date-time (e.g., 2026-01-19T12:00:00Z)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 20,
    "pattern": "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d{3})?Z?$",
    "validation": {
      "standard": "ISO 8601"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-infra--zone"></a>

### zone property

Type: `"string"`. Optional.

Instance zone (or region), depends on provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [infra.bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/bond_config/)
- [infra.hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hugepages/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/hw_info/)
- [infra.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/interfaces/)
- [infra.internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/internet_proxy/)
- [infra.sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/infra/sw_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/properties/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/registration/)
