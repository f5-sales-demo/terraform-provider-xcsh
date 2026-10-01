---
page_title: "Property reference"
subcategory: "DNS"
description: "Property reference for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 88431, "body_sha256": "sha256:843659e0e7734b083278427a4412dd9d1cd2ae8032b76b4454d0bfbfc52154e8", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary", "xcsh-docs:data-sources:dns_zone:properties:secondary"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:reference", "parent_id": "xcsh-docs:data-sources:dns_zone:fundamentals", "path": "documentation/data-sources/dns_zone/properties/index.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- Property reference

## Direct properties

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the DNSZone.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the DNSZone.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace where the DNSZone exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/): complete subsection reference.

- [secondary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/#schema-annotations) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/#schema-id) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/#schema-labels) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/#schema-namespace) |
| `primary` | [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/#section) |
| `primary.allow_http_lb_managed_records` | [primary.allow_http_lb_managed_records](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/#schema-primary--allow_http_lb_managed_records) |
| `primary.default_rr_set_group` | [primary.default_rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/#section) |
| `primary.default_rr_set_group.a_record` | [primary.default_rr_set_group.a_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/a_record/#section) |
| `primary.default_rr_set_group.a_record.name` | [primary.default_rr_set_group.a_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/a_record/#schema-primary--default_rr_set_group--a_record--name) |
| `primary.default_rr_set_group.a_record.values` | [primary.default_rr_set_group.a_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/a_record/#schema-primary--default_rr_set_group--a_record--values) |
| `primary.default_rr_set_group.aaaa_record` | [primary.default_rr_set_group.aaaa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/aaaa_record/#section) |
| `primary.default_rr_set_group.aaaa_record.name` | [primary.default_rr_set_group.aaaa_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/aaaa_record/#schema-primary--default_rr_set_group--aaaa_record--name) |
| `primary.default_rr_set_group.aaaa_record.values` | [primary.default_rr_set_group.aaaa_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/aaaa_record/#schema-primary--default_rr_set_group--aaaa_record--values) |
| `primary.default_rr_set_group.afsdb_record` | [primary.default_rr_set_group.afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/#section) |
| `primary.default_rr_set_group.afsdb_record.name` | [primary.default_rr_set_group.afsdb_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/#schema-primary--default_rr_set_group--afsdb_record--name) |
| `primary.default_rr_set_group.afsdb_record.values` | [primary.default_rr_set_group.afsdb_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/values/#section) |
| `primary.default_rr_set_group.afsdb_record.values.hostname` | [primary.default_rr_set_group.afsdb_record.values.hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/values/#schema-primary--default_rr_set_group--afsdb_record--values--hostname) |
| `primary.default_rr_set_group.afsdb_record.values.subtype` | [primary.default_rr_set_group.afsdb_record.values.subtype](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/afsdb_record/values/#schema-primary--default_rr_set_group--afsdb_record--values--subtype) |
| `primary.default_rr_set_group.alias_record` | [primary.default_rr_set_group.alias_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/alias_record/#section) |
| `primary.default_rr_set_group.alias_record.value` | [primary.default_rr_set_group.alias_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/alias_record/#schema-primary--default_rr_set_group--alias_record--value) |
| `primary.default_rr_set_group.caa_record` | [primary.default_rr_set_group.caa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/caa_record/#section) |
| `primary.default_rr_set_group.caa_record.name` | [primary.default_rr_set_group.caa_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/caa_record/#schema-primary--default_rr_set_group--caa_record--name) |
| `primary.default_rr_set_group.caa_record.values` | [primary.default_rr_set_group.caa_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/caa_record/values/#section) |
| `primary.default_rr_set_group.caa_record.values.flags` | [primary.default_rr_set_group.caa_record.values.flags](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/caa_record/values/#schema-primary--default_rr_set_group--caa_record--values--flags) |
| `primary.default_rr_set_group.caa_record.values.tag` | [primary.default_rr_set_group.caa_record.values.tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/caa_record/values/#schema-primary--default_rr_set_group--caa_record--values--tag) |
| `primary.default_rr_set_group.caa_record.values.value` | [primary.default_rr_set_group.caa_record.values.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/caa_record/values/#schema-primary--default_rr_set_group--caa_record--values--value) |
| `primary.default_rr_set_group.cds_record` | [primary.default_rr_set_group.cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/#section) |
| `primary.default_rr_set_group.cds_record.name` | [primary.default_rr_set_group.cds_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/#schema-primary--default_rr_set_group--cds_record--name) |
| `primary.default_rr_set_group.cds_record.values` | [primary.default_rr_set_group.cds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/#section) |
| `primary.default_rr_set_group.cds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.cds_record.values.ds_key_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/#schema-primary--default_rr_set_group--cds_record--values--ds_key_algorithm) |
| `primary.default_rr_set_group.cds_record.values.key_tag` | [primary.default_rr_set_group.cds_record.values.key_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/#schema-primary--default_rr_set_group--cds_record--values--key_tag) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha1_digest/#section) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha1_digest/#schema-primary--default_rr_set_group--cds_record--values--sha1_digest--digest) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha256_digest/#section) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha256_digest/#schema-primary--default_rr_set_group--cds_record--values--sha256_digest--digest) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha384_digest/#section) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cds_record/values/sha384_digest/#schema-primary--default_rr_set_group--cds_record--values--sha384_digest--digest) |
| `primary.default_rr_set_group.cert_record` | [primary.default_rr_set_group.cert_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cert_record/#section) |
| `primary.default_rr_set_group.cert_record.name` | [primary.default_rr_set_group.cert_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cert_record/#schema-primary--default_rr_set_group--cert_record--name) |
| `primary.default_rr_set_group.cert_record.values` | [primary.default_rr_set_group.cert_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cert_record/values/#section) |
| `primary.default_rr_set_group.cert_record.values.algorithm` | [primary.default_rr_set_group.cert_record.values.algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cert_record/values/#schema-primary--default_rr_set_group--cert_record--values--algorithm) |
| `primary.default_rr_set_group.cert_record.values.cert_key_tag` | [primary.default_rr_set_group.cert_record.values.cert_key_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cert_record/values/#schema-primary--default_rr_set_group--cert_record--values--cert_key_tag) |
| `primary.default_rr_set_group.cert_record.values.cert_type` | [primary.default_rr_set_group.cert_record.values.cert_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cert_record/values/#schema-primary--default_rr_set_group--cert_record--values--cert_type) |
| `primary.default_rr_set_group.cert_record.values.certificate` | [primary.default_rr_set_group.cert_record.values.certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cert_record/values/#schema-primary--default_rr_set_group--cert_record--values--certificate) |
| `primary.default_rr_set_group.cname_record` | [primary.default_rr_set_group.cname_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cname_record/#section) |
| `primary.default_rr_set_group.cname_record.name` | [primary.default_rr_set_group.cname_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cname_record/#schema-primary--default_rr_set_group--cname_record--name) |
| `primary.default_rr_set_group.cname_record.value` | [primary.default_rr_set_group.cname_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/cname_record/#schema-primary--default_rr_set_group--cname_record--value) |
| `primary.default_rr_set_group.description_spec` | [primary.default_rr_set_group.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/#schema-primary--default_rr_set_group--description_spec) |
| `primary.default_rr_set_group.ds_record` | [primary.default_rr_set_group.ds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/#section) |
| `primary.default_rr_set_group.ds_record.name` | [primary.default_rr_set_group.ds_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/#schema-primary--default_rr_set_group--ds_record--name) |
| `primary.default_rr_set_group.ds_record.values` | [primary.default_rr_set_group.ds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/#section) |
| `primary.default_rr_set_group.ds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.ds_record.values.ds_key_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/#schema-primary--default_rr_set_group--ds_record--values--ds_key_algorithm) |
| `primary.default_rr_set_group.ds_record.values.key_tag` | [primary.default_rr_set_group.ds_record.values.key_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/#schema-primary--default_rr_set_group--ds_record--values--key_tag) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha1_digest/#section) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha1_digest/#schema-primary--default_rr_set_group--ds_record--values--sha1_digest--digest) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha256_digest/#section) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha256_digest/#schema-primary--default_rr_set_group--ds_record--values--sha256_digest--digest) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha384_digest/#section) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ds_record/values/sha384_digest/#schema-primary--default_rr_set_group--ds_record--values--sha384_digest--digest) |
| `primary.default_rr_set_group.eui48_record` | [primary.default_rr_set_group.eui48_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui48_record/#section) |
| `primary.default_rr_set_group.eui48_record.name` | [primary.default_rr_set_group.eui48_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui48_record/#schema-primary--default_rr_set_group--eui48_record--name) |
| `primary.default_rr_set_group.eui48_record.value` | [primary.default_rr_set_group.eui48_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui48_record/#schema-primary--default_rr_set_group--eui48_record--value) |
| `primary.default_rr_set_group.eui64_record` | [primary.default_rr_set_group.eui64_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui64_record/#section) |
| `primary.default_rr_set_group.eui64_record.name` | [primary.default_rr_set_group.eui64_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui64_record/#schema-primary--default_rr_set_group--eui64_record--name) |
| `primary.default_rr_set_group.eui64_record.value` | [primary.default_rr_set_group.eui64_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/eui64_record/#schema-primary--default_rr_set_group--eui64_record--value) |
| `primary.default_rr_set_group.lb_record` | [primary.default_rr_set_group.lb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/#section) |
| `primary.default_rr_set_group.lb_record.name` | [primary.default_rr_set_group.lb_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/#schema-primary--default_rr_set_group--lb_record--name) |
| `primary.default_rr_set_group.lb_record.value` | [primary.default_rr_set_group.lb_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/value/#section) |
| `primary.default_rr_set_group.lb_record.value.name` | [primary.default_rr_set_group.lb_record.value.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/value/#schema-primary--default_rr_set_group--lb_record--value--name) |
| `primary.default_rr_set_group.lb_record.value.namespace` | [primary.default_rr_set_group.lb_record.value.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/value/#schema-primary--default_rr_set_group--lb_record--value--namespace) |
| `primary.default_rr_set_group.lb_record.value.tenant` | [primary.default_rr_set_group.lb_record.value.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/lb_record/value/#schema-primary--default_rr_set_group--lb_record--value--tenant) |
| `primary.default_rr_set_group.loc_record` | [primary.default_rr_set_group.loc_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/#section) |
| `primary.default_rr_set_group.loc_record.name` | [primary.default_rr_set_group.loc_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/#schema-primary--default_rr_set_group--loc_record--name) |
| `primary.default_rr_set_group.loc_record.values` | [primary.default_rr_set_group.loc_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#section) |
| `primary.default_rr_set_group.loc_record.values.altitude` | [primary.default_rr_set_group.loc_record.values.altitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--altitude) |
| `primary.default_rr_set_group.loc_record.values.horizontal_precision` | [primary.default_rr_set_group.loc_record.values.horizontal_precision](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--horizontal_precision) |
| `primary.default_rr_set_group.loc_record.values.latitude_degree` | [primary.default_rr_set_group.loc_record.values.latitude_degree](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--latitude_degree) |
| `primary.default_rr_set_group.loc_record.values.latitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.latitude_hemisphere](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--latitude_hemisphere) |
| `primary.default_rr_set_group.loc_record.values.latitude_minute` | [primary.default_rr_set_group.loc_record.values.latitude_minute](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--latitude_minute) |
| `primary.default_rr_set_group.loc_record.values.latitude_second` | [primary.default_rr_set_group.loc_record.values.latitude_second](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--latitude_second) |
| `primary.default_rr_set_group.loc_record.values.location_diameter` | [primary.default_rr_set_group.loc_record.values.location_diameter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--location_diameter) |
| `primary.default_rr_set_group.loc_record.values.longitude_degree` | [primary.default_rr_set_group.loc_record.values.longitude_degree](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--longitude_degree) |
| `primary.default_rr_set_group.loc_record.values.longitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.longitude_hemisphere](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--longitude_hemisphere) |
| `primary.default_rr_set_group.loc_record.values.longitude_minute` | [primary.default_rr_set_group.loc_record.values.longitude_minute](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--longitude_minute) |
| `primary.default_rr_set_group.loc_record.values.longitude_second` | [primary.default_rr_set_group.loc_record.values.longitude_second](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--longitude_second) |
| `primary.default_rr_set_group.loc_record.values.vertical_precision` | [primary.default_rr_set_group.loc_record.values.vertical_precision](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/loc_record/values/#schema-primary--default_rr_set_group--loc_record--values--vertical_precision) |
| `primary.default_rr_set_group.mx_record` | [primary.default_rr_set_group.mx_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/mx_record/#section) |
| `primary.default_rr_set_group.mx_record.name` | [primary.default_rr_set_group.mx_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/mx_record/#schema-primary--default_rr_set_group--mx_record--name) |
| `primary.default_rr_set_group.mx_record.values` | [primary.default_rr_set_group.mx_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/mx_record/values/#section) |
| `primary.default_rr_set_group.mx_record.values.domain` | [primary.default_rr_set_group.mx_record.values.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/mx_record/values/#schema-primary--default_rr_set_group--mx_record--values--domain) |
| `primary.default_rr_set_group.mx_record.values.priority` | [primary.default_rr_set_group.mx_record.values.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/mx_record/values/#schema-primary--default_rr_set_group--mx_record--values--priority) |
| `primary.default_rr_set_group.naptr_record` | [primary.default_rr_set_group.naptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/#section) |
| `primary.default_rr_set_group.naptr_record.name` | [primary.default_rr_set_group.naptr_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/#schema-primary--default_rr_set_group--naptr_record--name) |
| `primary.default_rr_set_group.naptr_record.values` | [primary.default_rr_set_group.naptr_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/values/#section) |
| `primary.default_rr_set_group.naptr_record.values.flags` | [primary.default_rr_set_group.naptr_record.values.flags](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/values/#schema-primary--default_rr_set_group--naptr_record--values--flags) |
| `primary.default_rr_set_group.naptr_record.values.order` | [primary.default_rr_set_group.naptr_record.values.order](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/values/#schema-primary--default_rr_set_group--naptr_record--values--order) |
| `primary.default_rr_set_group.naptr_record.values.preference` | [primary.default_rr_set_group.naptr_record.values.preference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/values/#schema-primary--default_rr_set_group--naptr_record--values--preference) |
| `primary.default_rr_set_group.naptr_record.values.regexp` | [primary.default_rr_set_group.naptr_record.values.regexp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/values/#schema-primary--default_rr_set_group--naptr_record--values--regexp) |
| `primary.default_rr_set_group.naptr_record.values.replacement` | [primary.default_rr_set_group.naptr_record.values.replacement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/values/#schema-primary--default_rr_set_group--naptr_record--values--replacement) |
| `primary.default_rr_set_group.naptr_record.values.service` | [primary.default_rr_set_group.naptr_record.values.service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/naptr_record/values/#schema-primary--default_rr_set_group--naptr_record--values--service) |
| `primary.default_rr_set_group.ns_record` | [primary.default_rr_set_group.ns_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ns_record/#section) |
| `primary.default_rr_set_group.ns_record.name` | [primary.default_rr_set_group.ns_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ns_record/#schema-primary--default_rr_set_group--ns_record--name) |
| `primary.default_rr_set_group.ns_record.values` | [primary.default_rr_set_group.ns_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ns_record/#schema-primary--default_rr_set_group--ns_record--values) |
| `primary.default_rr_set_group.ptr_record` | [primary.default_rr_set_group.ptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ptr_record/#section) |
| `primary.default_rr_set_group.ptr_record.name` | [primary.default_rr_set_group.ptr_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ptr_record/#schema-primary--default_rr_set_group--ptr_record--name) |
| `primary.default_rr_set_group.ptr_record.values` | [primary.default_rr_set_group.ptr_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/ptr_record/#schema-primary--default_rr_set_group--ptr_record--values) |
| `primary.default_rr_set_group.srv_record` | [primary.default_rr_set_group.srv_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/srv_record/#section) |
| `primary.default_rr_set_group.srv_record.name` | [primary.default_rr_set_group.srv_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/srv_record/#schema-primary--default_rr_set_group--srv_record--name) |
| `primary.default_rr_set_group.srv_record.values` | [primary.default_rr_set_group.srv_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/srv_record/values/#section) |
| `primary.default_rr_set_group.srv_record.values.port` | [primary.default_rr_set_group.srv_record.values.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/srv_record/values/#schema-primary--default_rr_set_group--srv_record--values--port) |
| `primary.default_rr_set_group.srv_record.values.priority` | [primary.default_rr_set_group.srv_record.values.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/srv_record/values/#schema-primary--default_rr_set_group--srv_record--values--priority) |
| `primary.default_rr_set_group.srv_record.values.target` | [primary.default_rr_set_group.srv_record.values.target](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/srv_record/values/#schema-primary--default_rr_set_group--srv_record--values--target) |
| `primary.default_rr_set_group.srv_record.values.weight` | [primary.default_rr_set_group.srv_record.values.weight](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/srv_record/values/#schema-primary--default_rr_set_group--srv_record--values--weight) |
| `primary.default_rr_set_group.sshfp_record` | [primary.default_rr_set_group.sshfp_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/#section) |
| `primary.default_rr_set_group.sshfp_record.name` | [primary.default_rr_set_group.sshfp_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/#schema-primary--default_rr_set_group--sshfp_record--name) |
| `primary.default_rr_set_group.sshfp_record.values` | [primary.default_rr_set_group.sshfp_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/#section) |
| `primary.default_rr_set_group.sshfp_record.values.algorithm` | [primary.default_rr_set_group.sshfp_record.values.algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/#schema-primary--default_rr_set_group--sshfp_record--values--algorithm) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/sha1_fingerprint/#section) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/sha1_fingerprint/#schema-primary--default_rr_set_group--sshfp_record--values--sha1_fingerprint--fingerprint) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/sha256_fingerprint/#section) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/sshfp_record/values/sha256_fingerprint/#schema-primary--default_rr_set_group--sshfp_record--values--sha256_fingerprint--fingerprint) |
| `primary.default_rr_set_group.tlsa_record` | [primary.default_rr_set_group.tlsa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/tlsa_record/#section) |
| `primary.default_rr_set_group.tlsa_record.name` | [primary.default_rr_set_group.tlsa_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/tlsa_record/#schema-primary--default_rr_set_group--tlsa_record--name) |
| `primary.default_rr_set_group.tlsa_record.values` | [primary.default_rr_set_group.tlsa_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/tlsa_record/values/#section) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_association_data` | [primary.default_rr_set_group.tlsa_record.values.certificate_association_data](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/tlsa_record/values/#schema-primary--default_rr_set_group--tlsa_record--values--certificate_association_data) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_usage` | [primary.default_rr_set_group.tlsa_record.values.certificate_usage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/tlsa_record/values/#schema-primary--default_rr_set_group--tlsa_record--values--certificate_usage) |
| `primary.default_rr_set_group.tlsa_record.values.matching_type` | [primary.default_rr_set_group.tlsa_record.values.matching_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/tlsa_record/values/#schema-primary--default_rr_set_group--tlsa_record--values--matching_type) |
| `primary.default_rr_set_group.tlsa_record.values.selector` | [primary.default_rr_set_group.tlsa_record.values.selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/tlsa_record/values/#schema-primary--default_rr_set_group--tlsa_record--values--selector) |
| `primary.default_rr_set_group.ttl` | [primary.default_rr_set_group.ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/#schema-primary--default_rr_set_group--ttl) |
| `primary.default_rr_set_group.txt_record` | [primary.default_rr_set_group.txt_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/txt_record/#section) |
| `primary.default_rr_set_group.txt_record.name` | [primary.default_rr_set_group.txt_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/txt_record/#schema-primary--default_rr_set_group--txt_record--name) |
| `primary.default_rr_set_group.txt_record.values` | [primary.default_rr_set_group.txt_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_rr_set_group/txt_record/#schema-primary--default_rr_set_group--txt_record--values) |
| `primary.default_soa_parameters` | [primary.default_soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/default_soa_parameters/#section) |
| `primary.dnssec_mode` | [primary.dnssec_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/dnssec_mode/#section) |
| `primary.dnssec_mode.disable_spec` | [primary.dnssec_mode.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/dnssec_mode/disable_spec/#section) |
| `primary.dnssec_mode.enable` | [primary.dnssec_mode.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/dnssec_mode/enable/#section) |
| `primary.rr_set_group` | [primary.rr_set_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/#section) |
| `primary.rr_set_group.metadata` | [primary.rr_set_group.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/metadata/#section) |
| `primary.rr_set_group.metadata.description_spec` | [primary.rr_set_group.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/metadata/#schema-primary--rr_set_group--metadata--description_spec) |
| `primary.rr_set_group.metadata.name` | [primary.rr_set_group.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/metadata/#schema-primary--rr_set_group--metadata--name) |
| `primary.rr_set_group.rr_set` | [primary.rr_set_group.rr_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/#section) |
| `primary.rr_set_group.rr_set.a_record` | [primary.rr_set_group.rr_set.a_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/a_record/#section) |
| `primary.rr_set_group.rr_set.a_record.name` | [primary.rr_set_group.rr_set.a_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/a_record/#schema-primary--rr_set_group--rr_set--a_record--name) |
| `primary.rr_set_group.rr_set.a_record.values` | [primary.rr_set_group.rr_set.a_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/a_record/#schema-primary--rr_set_group--rr_set--a_record--values) |
| `primary.rr_set_group.rr_set.aaaa_record` | [primary.rr_set_group.rr_set.aaaa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/aaaa_record/#section) |
| `primary.rr_set_group.rr_set.aaaa_record.name` | [primary.rr_set_group.rr_set.aaaa_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/aaaa_record/#schema-primary--rr_set_group--rr_set--aaaa_record--name) |
| `primary.rr_set_group.rr_set.aaaa_record.values` | [primary.rr_set_group.rr_set.aaaa_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/aaaa_record/#schema-primary--rr_set_group--rr_set--aaaa_record--values) |
| `primary.rr_set_group.rr_set.afsdb_record` | [primary.rr_set_group.rr_set.afsdb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/#section) |
| `primary.rr_set_group.rr_set.afsdb_record.name` | [primary.rr_set_group.rr_set.afsdb_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/#schema-primary--rr_set_group--rr_set--afsdb_record--name) |
| `primary.rr_set_group.rr_set.afsdb_record.values` | [primary.rr_set_group.rr_set.afsdb_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/values/#section) |
| `primary.rr_set_group.rr_set.afsdb_record.values.hostname` | [primary.rr_set_group.rr_set.afsdb_record.values.hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/values/#schema-primary--rr_set_group--rr_set--afsdb_record--values--hostname) |
| `primary.rr_set_group.rr_set.afsdb_record.values.subtype` | [primary.rr_set_group.rr_set.afsdb_record.values.subtype](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/afsdb_record/values/#schema-primary--rr_set_group--rr_set--afsdb_record--values--subtype) |
| `primary.rr_set_group.rr_set.alias_record` | [primary.rr_set_group.rr_set.alias_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/alias_record/#section) |
| `primary.rr_set_group.rr_set.alias_record.value` | [primary.rr_set_group.rr_set.alias_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/alias_record/#schema-primary--rr_set_group--rr_set--alias_record--value) |
| `primary.rr_set_group.rr_set.caa_record` | [primary.rr_set_group.rr_set.caa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/caa_record/#section) |
| `primary.rr_set_group.rr_set.caa_record.name` | [primary.rr_set_group.rr_set.caa_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/caa_record/#schema-primary--rr_set_group--rr_set--caa_record--name) |
| `primary.rr_set_group.rr_set.caa_record.values` | [primary.rr_set_group.rr_set.caa_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/caa_record/values/#section) |
| `primary.rr_set_group.rr_set.caa_record.values.flags` | [primary.rr_set_group.rr_set.caa_record.values.flags](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/caa_record/values/#schema-primary--rr_set_group--rr_set--caa_record--values--flags) |
| `primary.rr_set_group.rr_set.caa_record.values.tag` | [primary.rr_set_group.rr_set.caa_record.values.tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/caa_record/values/#schema-primary--rr_set_group--rr_set--caa_record--values--tag) |
| `primary.rr_set_group.rr_set.caa_record.values.value` | [primary.rr_set_group.rr_set.caa_record.values.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/caa_record/values/#schema-primary--rr_set_group--rr_set--caa_record--values--value) |
| `primary.rr_set_group.rr_set.cds_record` | [primary.rr_set_group.rr_set.cds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/#section) |
| `primary.rr_set_group.rr_set.cds_record.name` | [primary.rr_set_group.rr_set.cds_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/#schema-primary--rr_set_group--rr_set--cds_record--name) |
| `primary.rr_set_group.rr_set.cds_record.values` | [primary.rr_set_group.rr_set.cds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/#section) |
| `primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/#schema-primary--rr_set_group--rr_set--cds_record--values--ds_key_algorithm) |
| `primary.rr_set_group.rr_set.cds_record.values.key_tag` | [primary.rr_set_group.rr_set.cds_record.values.key_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/#schema-primary--rr_set_group--rr_set--cds_record--values--key_tag) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha1_digest/#section) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha1_digest/#schema-primary--rr_set_group--rr_set--cds_record--values--sha1_digest--digest) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha256_digest/#section) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha256_digest/#schema-primary--rr_set_group--rr_set--cds_record--values--sha256_digest--digest) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha384_digest/#section) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cds_record/values/sha384_digest/#schema-primary--rr_set_group--rr_set--cds_record--values--sha384_digest--digest) |
| `primary.rr_set_group.rr_set.cert_record` | [primary.rr_set_group.rr_set.cert_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cert_record/#section) |
| `primary.rr_set_group.rr_set.cert_record.name` | [primary.rr_set_group.rr_set.cert_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cert_record/#schema-primary--rr_set_group--rr_set--cert_record--name) |
| `primary.rr_set_group.rr_set.cert_record.values` | [primary.rr_set_group.rr_set.cert_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cert_record/values/#section) |
| `primary.rr_set_group.rr_set.cert_record.values.algorithm` | [primary.rr_set_group.rr_set.cert_record.values.algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cert_record/values/#schema-primary--rr_set_group--rr_set--cert_record--values--algorithm) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_key_tag` | [primary.rr_set_group.rr_set.cert_record.values.cert_key_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cert_record/values/#schema-primary--rr_set_group--rr_set--cert_record--values--cert_key_tag) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_type` | [primary.rr_set_group.rr_set.cert_record.values.cert_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cert_record/values/#schema-primary--rr_set_group--rr_set--cert_record--values--cert_type) |
| `primary.rr_set_group.rr_set.cert_record.values.certificate` | [primary.rr_set_group.rr_set.cert_record.values.certificate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cert_record/values/#schema-primary--rr_set_group--rr_set--cert_record--values--certificate) |
| `primary.rr_set_group.rr_set.cname_record` | [primary.rr_set_group.rr_set.cname_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cname_record/#section) |
| `primary.rr_set_group.rr_set.cname_record.name` | [primary.rr_set_group.rr_set.cname_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cname_record/#schema-primary--rr_set_group--rr_set--cname_record--name) |
| `primary.rr_set_group.rr_set.cname_record.value` | [primary.rr_set_group.rr_set.cname_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/cname_record/#schema-primary--rr_set_group--rr_set--cname_record--value) |
| `primary.rr_set_group.rr_set.description_spec` | [primary.rr_set_group.rr_set.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/#schema-primary--rr_set_group--rr_set--description_spec) |
| `primary.rr_set_group.rr_set.ds_record` | [primary.rr_set_group.rr_set.ds_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/#section) |
| `primary.rr_set_group.rr_set.ds_record.name` | [primary.rr_set_group.rr_set.ds_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/#schema-primary--rr_set_group--rr_set--ds_record--name) |
| `primary.rr_set_group.rr_set.ds_record.values` | [primary.rr_set_group.rr_set.ds_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/#section) |
| `primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/#schema-primary--rr_set_group--rr_set--ds_record--values--ds_key_algorithm) |
| `primary.rr_set_group.rr_set.ds_record.values.key_tag` | [primary.rr_set_group.rr_set.ds_record.values.key_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/#schema-primary--rr_set_group--rr_set--ds_record--values--key_tag) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/sha1_digest/#section) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/sha1_digest/#schema-primary--rr_set_group--rr_set--ds_record--values--sha1_digest--digest) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/sha256_digest/#section) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/sha256_digest/#schema-primary--rr_set_group--rr_set--ds_record--values--sha256_digest--digest) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/sha384_digest/#section) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ds_record/values/sha384_digest/#schema-primary--rr_set_group--rr_set--ds_record--values--sha384_digest--digest) |
| `primary.rr_set_group.rr_set.eui48_record` | [primary.rr_set_group.rr_set.eui48_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui48_record/#section) |
| `primary.rr_set_group.rr_set.eui48_record.name` | [primary.rr_set_group.rr_set.eui48_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui48_record/#schema-primary--rr_set_group--rr_set--eui48_record--name) |
| `primary.rr_set_group.rr_set.eui48_record.value` | [primary.rr_set_group.rr_set.eui48_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui48_record/#schema-primary--rr_set_group--rr_set--eui48_record--value) |
| `primary.rr_set_group.rr_set.eui64_record` | [primary.rr_set_group.rr_set.eui64_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui64_record/#section) |
| `primary.rr_set_group.rr_set.eui64_record.name` | [primary.rr_set_group.rr_set.eui64_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui64_record/#schema-primary--rr_set_group--rr_set--eui64_record--name) |
| `primary.rr_set_group.rr_set.eui64_record.value` | [primary.rr_set_group.rr_set.eui64_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/eui64_record/#schema-primary--rr_set_group--rr_set--eui64_record--value) |
| `primary.rr_set_group.rr_set.lb_record` | [primary.rr_set_group.rr_set.lb_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/lb_record/#section) |
| `primary.rr_set_group.rr_set.lb_record.name` | [primary.rr_set_group.rr_set.lb_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/lb_record/#schema-primary--rr_set_group--rr_set--lb_record--name) |
| `primary.rr_set_group.rr_set.lb_record.value` | [primary.rr_set_group.rr_set.lb_record.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/lb_record/value/#section) |
| `primary.rr_set_group.rr_set.lb_record.value.name` | [primary.rr_set_group.rr_set.lb_record.value.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/lb_record/value/#schema-primary--rr_set_group--rr_set--lb_record--value--name) |
| `primary.rr_set_group.rr_set.lb_record.value.namespace` | [primary.rr_set_group.rr_set.lb_record.value.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/lb_record/value/#schema-primary--rr_set_group--rr_set--lb_record--value--namespace) |
| `primary.rr_set_group.rr_set.lb_record.value.tenant` | [primary.rr_set_group.rr_set.lb_record.value.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/lb_record/value/#schema-primary--rr_set_group--rr_set--lb_record--value--tenant) |
| `primary.rr_set_group.rr_set.loc_record` | [primary.rr_set_group.rr_set.loc_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/#section) |
| `primary.rr_set_group.rr_set.loc_record.name` | [primary.rr_set_group.rr_set.loc_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/#schema-primary--rr_set_group--rr_set--loc_record--name) |
| `primary.rr_set_group.rr_set.loc_record.values` | [primary.rr_set_group.rr_set.loc_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#section) |
| `primary.rr_set_group.rr_set.loc_record.values.altitude` | [primary.rr_set_group.rr_set.loc_record.values.altitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--altitude) |
| `primary.rr_set_group.rr_set.loc_record.values.horizontal_precision` | [primary.rr_set_group.rr_set.loc_record.values.horizontal_precision](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--horizontal_precision) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.latitude_degree](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--latitude_degree) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--latitude_hemisphere) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.latitude_minute](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--latitude_minute) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_second` | [primary.rr_set_group.rr_set.loc_record.values.latitude_second](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--latitude_second) |
| `primary.rr_set_group.rr_set.loc_record.values.location_diameter` | [primary.rr_set_group.rr_set.loc_record.values.location_diameter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--location_diameter) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.longitude_degree](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--longitude_degree) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--longitude_hemisphere) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.longitude_minute](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--longitude_minute) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_second` | [primary.rr_set_group.rr_set.loc_record.values.longitude_second](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--longitude_second) |
| `primary.rr_set_group.rr_set.loc_record.values.vertical_precision` | [primary.rr_set_group.rr_set.loc_record.values.vertical_precision](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/loc_record/values/#schema-primary--rr_set_group--rr_set--loc_record--values--vertical_precision) |
| `primary.rr_set_group.rr_set.mx_record` | [primary.rr_set_group.rr_set.mx_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/mx_record/#section) |
| `primary.rr_set_group.rr_set.mx_record.name` | [primary.rr_set_group.rr_set.mx_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/mx_record/#schema-primary--rr_set_group--rr_set--mx_record--name) |
| `primary.rr_set_group.rr_set.mx_record.values` | [primary.rr_set_group.rr_set.mx_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/mx_record/values/#section) |
| `primary.rr_set_group.rr_set.mx_record.values.domain` | [primary.rr_set_group.rr_set.mx_record.values.domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/mx_record/values/#schema-primary--rr_set_group--rr_set--mx_record--values--domain) |
| `primary.rr_set_group.rr_set.mx_record.values.priority` | [primary.rr_set_group.rr_set.mx_record.values.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/mx_record/values/#schema-primary--rr_set_group--rr_set--mx_record--values--priority) |
| `primary.rr_set_group.rr_set.naptr_record` | [primary.rr_set_group.rr_set.naptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/#section) |
| `primary.rr_set_group.rr_set.naptr_record.name` | [primary.rr_set_group.rr_set.naptr_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/#schema-primary--rr_set_group--rr_set--naptr_record--name) |
| `primary.rr_set_group.rr_set.naptr_record.values` | [primary.rr_set_group.rr_set.naptr_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/#section) |
| `primary.rr_set_group.rr_set.naptr_record.values.flags` | [primary.rr_set_group.rr_set.naptr_record.values.flags](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/#schema-primary--rr_set_group--rr_set--naptr_record--values--flags) |
| `primary.rr_set_group.rr_set.naptr_record.values.order` | [primary.rr_set_group.rr_set.naptr_record.values.order](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/#schema-primary--rr_set_group--rr_set--naptr_record--values--order) |
| `primary.rr_set_group.rr_set.naptr_record.values.preference` | [primary.rr_set_group.rr_set.naptr_record.values.preference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/#schema-primary--rr_set_group--rr_set--naptr_record--values--preference) |
| `primary.rr_set_group.rr_set.naptr_record.values.regexp` | [primary.rr_set_group.rr_set.naptr_record.values.regexp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/#schema-primary--rr_set_group--rr_set--naptr_record--values--regexp) |
| `primary.rr_set_group.rr_set.naptr_record.values.replacement` | [primary.rr_set_group.rr_set.naptr_record.values.replacement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/#schema-primary--rr_set_group--rr_set--naptr_record--values--replacement) |
| `primary.rr_set_group.rr_set.naptr_record.values.service` | [primary.rr_set_group.rr_set.naptr_record.values.service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/naptr_record/values/#schema-primary--rr_set_group--rr_set--naptr_record--values--service) |
| `primary.rr_set_group.rr_set.ns_record` | [primary.rr_set_group.rr_set.ns_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ns_record/#section) |
| `primary.rr_set_group.rr_set.ns_record.name` | [primary.rr_set_group.rr_set.ns_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ns_record/#schema-primary--rr_set_group--rr_set--ns_record--name) |
| `primary.rr_set_group.rr_set.ns_record.values` | [primary.rr_set_group.rr_set.ns_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ns_record/#schema-primary--rr_set_group--rr_set--ns_record--values) |
| `primary.rr_set_group.rr_set.ptr_record` | [primary.rr_set_group.rr_set.ptr_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ptr_record/#section) |
| `primary.rr_set_group.rr_set.ptr_record.name` | [primary.rr_set_group.rr_set.ptr_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ptr_record/#schema-primary--rr_set_group--rr_set--ptr_record--name) |
| `primary.rr_set_group.rr_set.ptr_record.values` | [primary.rr_set_group.rr_set.ptr_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/ptr_record/#schema-primary--rr_set_group--rr_set--ptr_record--values) |
| `primary.rr_set_group.rr_set.srv_record` | [primary.rr_set_group.rr_set.srv_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/srv_record/#section) |
| `primary.rr_set_group.rr_set.srv_record.name` | [primary.rr_set_group.rr_set.srv_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/srv_record/#schema-primary--rr_set_group--rr_set--srv_record--name) |
| `primary.rr_set_group.rr_set.srv_record.values` | [primary.rr_set_group.rr_set.srv_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/srv_record/values/#section) |
| `primary.rr_set_group.rr_set.srv_record.values.port` | [primary.rr_set_group.rr_set.srv_record.values.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/srv_record/values/#schema-primary--rr_set_group--rr_set--srv_record--values--port) |
| `primary.rr_set_group.rr_set.srv_record.values.priority` | [primary.rr_set_group.rr_set.srv_record.values.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/srv_record/values/#schema-primary--rr_set_group--rr_set--srv_record--values--priority) |
| `primary.rr_set_group.rr_set.srv_record.values.target` | [primary.rr_set_group.rr_set.srv_record.values.target](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/srv_record/values/#schema-primary--rr_set_group--rr_set--srv_record--values--target) |
| `primary.rr_set_group.rr_set.srv_record.values.weight` | [primary.rr_set_group.rr_set.srv_record.values.weight](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/srv_record/values/#schema-primary--rr_set_group--rr_set--srv_record--values--weight) |
| `primary.rr_set_group.rr_set.sshfp_record` | [primary.rr_set_group.rr_set.sshfp_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/#section) |
| `primary.rr_set_group.rr_set.sshfp_record.name` | [primary.rr_set_group.rr_set.sshfp_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/#schema-primary--rr_set_group--rr_set--sshfp_record--name) |
| `primary.rr_set_group.rr_set.sshfp_record.values` | [primary.rr_set_group.rr_set.sshfp_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/#section) |
| `primary.rr_set_group.rr_set.sshfp_record.values.algorithm` | [primary.rr_set_group.rr_set.sshfp_record.values.algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/#schema-primary--rr_set_group--rr_set--sshfp_record--values--algorithm) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/sha1_fingerprint/#section) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/sha1_fingerprint/#schema-primary--rr_set_group--rr_set--sshfp_record--values--sha1_fingerprint--fingerprint) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/sha256_fingerprint/#section) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/sshfp_record/values/sha256_fingerprint/#schema-primary--rr_set_group--rr_set--sshfp_record--values--sha256_fingerprint--fingerprint) |
| `primary.rr_set_group.rr_set.tlsa_record` | [primary.rr_set_group.rr_set.tlsa_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/#section) |
| `primary.rr_set_group.rr_set.tlsa_record.name` | [primary.rr_set_group.rr_set.tlsa_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/#schema-primary--rr_set_group--rr_set--tlsa_record--name) |
| `primary.rr_set_group.rr_set.tlsa_record.values` | [primary.rr_set_group.rr_set.tlsa_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/values/#section) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/values/#schema-primary--rr_set_group--rr_set--tlsa_record--values--certificate_association_data) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/values/#schema-primary--rr_set_group--rr_set--tlsa_record--values--certificate_usage) |
| `primary.rr_set_group.rr_set.tlsa_record.values.matching_type` | [primary.rr_set_group.rr_set.tlsa_record.values.matching_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/values/#schema-primary--rr_set_group--rr_set--tlsa_record--values--matching_type) |
| `primary.rr_set_group.rr_set.tlsa_record.values.selector` | [primary.rr_set_group.rr_set.tlsa_record.values.selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/tlsa_record/values/#schema-primary--rr_set_group--rr_set--tlsa_record--values--selector) |
| `primary.rr_set_group.rr_set.ttl` | [primary.rr_set_group.rr_set.ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/#schema-primary--rr_set_group--rr_set--ttl) |
| `primary.rr_set_group.rr_set.txt_record` | [primary.rr_set_group.rr_set.txt_record](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/txt_record/#section) |
| `primary.rr_set_group.rr_set.txt_record.name` | [primary.rr_set_group.rr_set.txt_record.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/txt_record/#schema-primary--rr_set_group--rr_set--txt_record--name) |
| `primary.rr_set_group.rr_set.txt_record.values` | [primary.rr_set_group.rr_set.txt_record.values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/rr_set_group/rr_set/txt_record/#schema-primary--rr_set_group--rr_set--txt_record--values) |
| `primary.soa_parameters` | [primary.soa_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/soa_parameters/#section) |
| `primary.soa_parameters.expire` | [primary.soa_parameters.expire](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/soa_parameters/#schema-primary--soa_parameters--expire) |
| `primary.soa_parameters.negative_ttl` | [primary.soa_parameters.negative_ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/soa_parameters/#schema-primary--soa_parameters--negative_ttl) |
| `primary.soa_parameters.refresh` | [primary.soa_parameters.refresh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/soa_parameters/#schema-primary--soa_parameters--refresh) |
| `primary.soa_parameters.retry` | [primary.soa_parameters.retry](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/soa_parameters/#schema-primary--soa_parameters--retry) |
| `primary.soa_parameters.ttl` | [primary.soa_parameters.ttl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/soa_parameters/#schema-primary--soa_parameters--ttl) |
| `secondary` | [secondary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/#section) |
| `secondary.primary_servers` | [secondary.primary_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/#schema-secondary--primary_servers) |
| `secondary.tsig_key_algorithm` | [secondary.tsig_key_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/#schema-secondary--tsig_key_algorithm) |
| `secondary.tsig_key_name` | [secondary.tsig_key_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/#schema-secondary--tsig_key_name) |
| `secondary.tsig_key_value` | [secondary.tsig_key_value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/#section) |
| `secondary.tsig_key_value.blindfold_secret_info` | [secondary.tsig_key_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/blindfold_secret_info/#section) |
| `secondary.tsig_key_value.blindfold_secret_info.decryption_provider` | [secondary.tsig_key_value.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/blindfold_secret_info/#schema-secondary--tsig_key_value--blindfold_secret_info--decryption_provider) |
| `secondary.tsig_key_value.blindfold_secret_info.location` | [secondary.tsig_key_value.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/blindfold_secret_info/#schema-secondary--tsig_key_value--blindfold_secret_info--location) |
| `secondary.tsig_key_value.blindfold_secret_info.store_provider` | [secondary.tsig_key_value.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/blindfold_secret_info/#schema-secondary--tsig_key_value--blindfold_secret_info--store_provider) |
| `secondary.tsig_key_value.clear_secret_info` | [secondary.tsig_key_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/clear_secret_info/#section) |
| `secondary.tsig_key_value.clear_secret_info.provider_ref` | [secondary.tsig_key_value.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/clear_secret_info/#schema-secondary--tsig_key_value--clear_secret_info--provider_ref) |
| `secondary.tsig_key_value.clear_secret_info.url` | [secondary.tsig_key_value.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/tsig_key_value/clear_secret_info/#schema-secondary--tsig_key_value--clear_secret_info--url) |

## Next pages

- [primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/primary/)
- [secondary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/secondary/)
- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
