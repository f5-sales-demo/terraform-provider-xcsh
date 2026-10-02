---
page_title: "infra"
subcategory: ""
description: "InfraMetadata stores information about instance infrastructure."
xcsh_docs: {"aliases": ["infra"], "body_bytes": 9853, "body_sha256": "sha256:314c0080899842ae5c8f17e3f7d85ba4fa88dd198280d6b243bdd4bd86137d23", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:registration:properties:infra:bond_config", "xcsh-docs:data-sources:registration:properties:infra:hugepages", "xcsh-docs:data-sources:registration:properties:infra:hw_info", "xcsh-docs:data-sources:registration:properties:infra:interfaces", "xcsh-docs:data-sources:registration:properties:infra:internet_proxy", "xcsh-docs:data-sources:registration:properties:infra:sw_info"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:registration:properties:infra", "parent_id": "xcsh-docs:data-sources:registration:reference", "path": "documentation/data-sources/registration/properties/infra/index.md", "product": "distributed-cloud", "provider_name": "registration", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1322220013303031-1121222131032033-0300200330111020-1211223110101013-1222330213123202-2313102202302131-0222130030200021-1110122111021101", "registry_path": "docs/guides/data-sources--registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["infra"], "schema_version": 1, "sections": [{"aliases": ["availability zone"], "anchor": "schema-infra--availability_zone", "description": "An Availability Zone is a high-availability offering that protects your applications and data from datacenter failures.", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "availability_zone"], "syntax": "attribute", "type": "string"}, {"aliases": ["bond config"], "anchor": "section", "description": "Bond device configuration for VPM registration.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:bond_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "bond_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["certified hw"], "anchor": "schema-infra--certified_hw", "description": "Certified HW name used to map with F5XC certified_hardware definition.", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["domain"], "anchor": "schema-infra--domain", "description": "Machine domain. It's used for Kubernetes cloud provider when domain must be different than F5 Distributed Cloud.", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["hostname"], "anchor": "schema-infra--hostname", "description": "Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["hugepages"], "anchor": "section", "description": "Hugepage settings for CE on K8s SMV2 site.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hugepages", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["infra", "hugepages"], "syntax": "attribute", "type": "object"}, {"aliases": ["hw info"], "anchor": "section", "description": "OsInfo holds information about host OS and HW.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:hw_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "hw_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["instance id"], "anchor": "schema-infra--instance_id", "description": "Instance ID (assigned by infrastructure provider)", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "instance_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["interfaces"], "anchor": "section", "description": "Machine interfaces present during registration time.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:interfaces", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "interfaces"], "syntax": "attribute", "type": "object"}, {"aliases": ["internet proxy"], "anchor": "section", "description": "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "internet_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["is slo static"], "anchor": "schema-infra--is_slo_static", "description": "Indicates whether the SLO is static.", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "is_slo_static"], "syntax": "attribute", "type": "bool"}, {"aliases": ["machine id"], "anchor": "schema-infra--machine_id", "description": "Machine ID - generated by operating system.", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "machine_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["provider ref"], "anchor": "schema-infra--provider_ref", "description": "Infrastructure provider enum for registration. It describes where is instance running. Provider was not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other provider, which was not identified by system. Possible values are `UNKNOWN`, `AWS`, `GOOGLE`, `AZURE`, `VMWARE`, `KVM`,", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "provider_ref"], "syntax": "attribute", "type": "string"}, {"aliases": ["sw info"], "anchor": "section", "description": "SWInfo holds information about sw version.", "document_id": "xcsh-docs:data-sources:registration:properties:infra:sw_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["infra", "sw_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["timestamp"], "anchor": "schema-infra--timestamp", "description": "It's used to verify machine have acceptable time difference from server.", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["zone"], "anchor": "schema-infra--zone", "description": "Instance zone (or region), depends on provider.", "document_id": "xcsh-docs:data-sources:registration:properties:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["infra", "zone"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/registration/properties/infra/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "InfraMetadata stores information about instance infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["registrationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# infra

Breadcrumbs:

- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- infra

<a id="section"></a>

Type: `"single"`. Computed.

InfraMetadata stores information about instance infrastructure.

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

## Direct properties

<a id="schema-infra--availability_zone"></a>

### availability_zone property

Type: `"string"`. Computed.

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

- [bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/bond_config/): complete subsection reference.

<a id="schema-infra--certified_hw"></a>

### certified_hw property

Type: `"string"`. Computed.

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

Type: `"string"`. Computed.

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

Type: `"string"`. Computed.

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

Upstream description:

Must be unique in entire cluster and same as OS settings. '.' (dots) are not allowed in hostname.

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

- [hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hugepages/): complete subsection reference.

- [hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/): complete subsection reference.

<a id="schema-infra--instance_id"></a>

### instance_id property

Type: `"string"`. Computed.

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

- [interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/interfaces/): complete subsection reference.

- [internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/internet_proxy/): complete subsection reference.

<a id="schema-infra--is_slo_static"></a>

### is_slo_static property

Type: `"bool"`. Computed.

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

Type: `"string"`. Computed.

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

Type: `"string"`. Computed.

\[Enum:
UNKNOWN|AWS|GOOGLE|AZURE|VMWARE|KVM|OTHER|VOLTERRA|IBMCLOUD|UNKNOWN\_K8S|AWS\_K8S|GCP\_K8S|AZURE\_K8S|VMWARE\_K8S|KVM\_K8S|OTHER\_K8S|VOLTERRA\_K8S|IBMCLOUD\_K8S|F5OS|RSERIES|OCI|NUTANIX|OPENSTACK|EQUINIX|OPENSHIFT\_VIRTUALIZATION|KUBERNETES\]
Infrastructure provider enum for registration. It describes where is instance running. Provider was
not detected AWS cloud instance Google cloud instance Azure cloud instance VMWare VM KVM VM Other
provider, which was not identified by system. Possible values are \`UNKNOWN\`, \`AWS\`, \`GOOGLE\`,
\`AZURE\`, \`VMWARE\`, \`KVM\`, \`OTHER\`, \`VOLTERRA\`, \`IBMCLOUD\`, \`UNKNOWN\_K8S\`,
\`AWS\_K8S\`, \`GCP\_K8S\`, \`AZURE\_K8S\`, \`VMWARE\_K8S\`, \`KVM\_K8S\`, \`OTHER\_K8S\`,
\`VOLTERRA\_K8S\`, \`IBMCLOUD\_K8S\`, \`F5OS\`, \`RSERIES\`, \`OCI\`, \`NUTANIX\`, \`OPENSTACK\`,
\`EQUINIX\`, \`OPENSHIFT\_VIRTUALIZATION\`, \`KUBERNETES\`.

- [sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/sw_info/): complete subsection reference.

<a id="schema-infra--timestamp"></a>

### timestamp property

Type: `"string"`. Computed.

It's used to verify machine have acceptable time difference from server.

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

Type: `"string"`. Computed.

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

- [infra.bond_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/bond_config/)
- [infra.hugepages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hugepages/)
- [infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/hw_info/)
- [infra.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/interfaces/)
- [infra.internet_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/internet_proxy/)
- [infra.sw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/infra/sw_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/properties/)
- [xcsh_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/registration/)
