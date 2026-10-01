---
page_title: "Property reference"
subcategory: "DNS"
description: "Property reference for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 74363, "body_sha256": "sha256:6e61e68cb362a44cc3a35b55d19c9268a74b7adace0dfe59be0b0fe775d84bd0", "canonical_id": "xcsh-docs:data-sources:dns_zone:reference", "child_ids": ["xcsh-docs:data-sources:dns_zone:properties:primary", "xcsh-docs:data-sources:dns_zone:properties:secondary"], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:reference", "parent_id": "xcsh-docs:data-sources:dns_zone:fundamentals", "path": "docs/guides/data-sources--dns_zone--reference.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
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

- [primary](data-sources--dns_zone--properties--primary.md): complete subsection reference.

- [secondary](data-sources--dns_zone--properties--secondary.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_zone--reference.md#schema-annotations) |
| `description` | [description](data-sources--dns_zone--reference.md#schema-description) |
| `id` | [id](data-sources--dns_zone--reference.md#schema-id) |
| `labels` | [labels](data-sources--dns_zone--reference.md#schema-labels) |
| `name` | [name](data-sources--dns_zone--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--dns_zone--reference.md#schema-namespace) |
| `primary` | [primary](data-sources--dns_zone--properties--primary.md#section) |
| `primary.allow_http_lb_managed_records` | [primary.allow_http_lb_managed_records](data-sources--dns_zone--properties--primary.md#schema-primary--allow_http_lb_managed_records) |
| `primary.default_rr_set_group` | [primary.default_rr_set_group](data-sources--dns_zone--properties--primary--default_rr_set_group.md#section) |
| `primary.default_rr_set_group.a_record` | [primary.default_rr_set_group.a_record](data-sources--dns_zone--properties--primary--default_rr_set_group--a_record.md#section) |
| `primary.default_rr_set_group.a_record.name` | [primary.default_rr_set_group.a_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--a_record.md#schema-primary--default_rr_set_group--a_record--name) |
| `primary.default_rr_set_group.a_record.values` | [primary.default_rr_set_group.a_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--a_record.md#schema-primary--default_rr_set_group--a_record--values) |
| `primary.default_rr_set_group.aaaa_record` | [primary.default_rr_set_group.aaaa_record](data-sources--dns_zone--properties--primary--default_rr_set_group--aaaa_record.md#section) |
| `primary.default_rr_set_group.aaaa_record.name` | [primary.default_rr_set_group.aaaa_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--aaaa_record.md#schema-primary--default_rr_set_group--aaaa_record--name) |
| `primary.default_rr_set_group.aaaa_record.values` | [primary.default_rr_set_group.aaaa_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--aaaa_record.md#schema-primary--default_rr_set_group--aaaa_record--values) |
| `primary.default_rr_set_group.afsdb_record` | [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--properties--primary--default_rr_set_group--afsdb_record.md#section) |
| `primary.default_rr_set_group.afsdb_record.name` | [primary.default_rr_set_group.afsdb_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--afsdb_record.md#schema-primary--default_rr_set_group--afsdb_record--name) |
| `primary.default_rr_set_group.afsdb_record.values` | [primary.default_rr_set_group.afsdb_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--afsdb_record--values.md#section) |
| `primary.default_rr_set_group.afsdb_record.values.hostname` | [primary.default_rr_set_group.afsdb_record.values.hostname](data-sources--dns_zone--properties--primary--default_rr_set_group--afsdb_record--values.md#schema-primary--default_rr_set_group--afsdb_record--values--hostname) |
| `primary.default_rr_set_group.afsdb_record.values.subtype` | [primary.default_rr_set_group.afsdb_record.values.subtype](data-sources--dns_zone--properties--primary--default_rr_set_group--afsdb_record--values.md#schema-primary--default_rr_set_group--afsdb_record--values--subtype) |
| `primary.default_rr_set_group.alias_record` | [primary.default_rr_set_group.alias_record](data-sources--dns_zone--properties--primary--default_rr_set_group--alias_record.md#section) |
| `primary.default_rr_set_group.alias_record.value` | [primary.default_rr_set_group.alias_record.value](data-sources--dns_zone--properties--primary--default_rr_set_group--alias_record.md#schema-primary--default_rr_set_group--alias_record--value) |
| `primary.default_rr_set_group.caa_record` | [primary.default_rr_set_group.caa_record](data-sources--dns_zone--properties--primary--default_rr_set_group--caa_record.md#section) |
| `primary.default_rr_set_group.caa_record.name` | [primary.default_rr_set_group.caa_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--caa_record.md#schema-primary--default_rr_set_group--caa_record--name) |
| `primary.default_rr_set_group.caa_record.values` | [primary.default_rr_set_group.caa_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--caa_record--values.md#section) |
| `primary.default_rr_set_group.caa_record.values.flags` | [primary.default_rr_set_group.caa_record.values.flags](data-sources--dns_zone--properties--primary--default_rr_set_group--caa_record--values.md#schema-primary--default_rr_set_group--caa_record--values--flags) |
| `primary.default_rr_set_group.caa_record.values.tag` | [primary.default_rr_set_group.caa_record.values.tag](data-sources--dns_zone--properties--primary--default_rr_set_group--caa_record--values.md#schema-primary--default_rr_set_group--caa_record--values--tag) |
| `primary.default_rr_set_group.caa_record.values.value` | [primary.default_rr_set_group.caa_record.values.value](data-sources--dns_zone--properties--primary--default_rr_set_group--caa_record--values.md#schema-primary--default_rr_set_group--caa_record--values--value) |
| `primary.default_rr_set_group.cds_record` | [primary.default_rr_set_group.cds_record](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record.md#section) |
| `primary.default_rr_set_group.cds_record.name` | [primary.default_rr_set_group.cds_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record.md#schema-primary--default_rr_set_group--cds_record--name) |
| `primary.default_rr_set_group.cds_record.values` | [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record--values.md#section) |
| `primary.default_rr_set_group.cds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.cds_record.values.ds_key_algorithm](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record--values.md#schema-primary--default_rr_set_group--cds_record--values--ds_key_algorithm) |
| `primary.default_rr_set_group.cds_record.values.key_tag` | [primary.default_rr_set_group.cds_record.values.key_tag](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record--values.md#schema-primary--default_rr_set_group--cds_record--values--key_tag) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record--values--sha1_digest.md#section) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest.digest](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record--values--sha1_digest.md#schema-primary--default_rr_set_group--cds_record--values--sha1_digest--digest) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record--values--sha256_digest.md#section) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest.digest](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record--values--sha256_digest.md#schema-primary--default_rr_set_group--cds_record--values--sha256_digest--digest) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record--values--sha384_digest.md#section) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest.digest](data-sources--dns_zone--properties--primary--default_rr_set_group--cds_record--values--sha384_digest.md#schema-primary--default_rr_set_group--cds_record--values--sha384_digest--digest) |
| `primary.default_rr_set_group.cert_record` | [primary.default_rr_set_group.cert_record](data-sources--dns_zone--properties--primary--default_rr_set_group--cert_record.md#section) |
| `primary.default_rr_set_group.cert_record.name` | [primary.default_rr_set_group.cert_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--cert_record.md#schema-primary--default_rr_set_group--cert_record--name) |
| `primary.default_rr_set_group.cert_record.values` | [primary.default_rr_set_group.cert_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--cert_record--values.md#section) |
| `primary.default_rr_set_group.cert_record.values.algorithm` | [primary.default_rr_set_group.cert_record.values.algorithm](data-sources--dns_zone--properties--primary--default_rr_set_group--cert_record--values.md#schema-primary--default_rr_set_group--cert_record--values--algorithm) |
| `primary.default_rr_set_group.cert_record.values.cert_key_tag` | [primary.default_rr_set_group.cert_record.values.cert_key_tag](data-sources--dns_zone--properties--primary--default_rr_set_group--cert_record--values.md#schema-primary--default_rr_set_group--cert_record--values--cert_key_tag) |
| `primary.default_rr_set_group.cert_record.values.cert_type` | [primary.default_rr_set_group.cert_record.values.cert_type](data-sources--dns_zone--properties--primary--default_rr_set_group--cert_record--values.md#schema-primary--default_rr_set_group--cert_record--values--cert_type) |
| `primary.default_rr_set_group.cert_record.values.certificate` | [primary.default_rr_set_group.cert_record.values.certificate](data-sources--dns_zone--properties--primary--default_rr_set_group--cert_record--values.md#schema-primary--default_rr_set_group--cert_record--values--certificate) |
| `primary.default_rr_set_group.cname_record` | [primary.default_rr_set_group.cname_record](data-sources--dns_zone--properties--primary--default_rr_set_group--cname_record.md#section) |
| `primary.default_rr_set_group.cname_record.name` | [primary.default_rr_set_group.cname_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--cname_record.md#schema-primary--default_rr_set_group--cname_record--name) |
| `primary.default_rr_set_group.cname_record.value` | [primary.default_rr_set_group.cname_record.value](data-sources--dns_zone--properties--primary--default_rr_set_group--cname_record.md#schema-primary--default_rr_set_group--cname_record--value) |
| `primary.default_rr_set_group.description_spec` | [primary.default_rr_set_group.description_spec](data-sources--dns_zone--properties--primary--default_rr_set_group.md#schema-primary--default_rr_set_group--description_spec) |
| `primary.default_rr_set_group.ds_record` | [primary.default_rr_set_group.ds_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record.md#section) |
| `primary.default_rr_set_group.ds_record.name` | [primary.default_rr_set_group.ds_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record.md#schema-primary--default_rr_set_group--ds_record--name) |
| `primary.default_rr_set_group.ds_record.values` | [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values.md#section) |
| `primary.default_rr_set_group.ds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.ds_record.values.ds_key_algorithm](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values.md#schema-primary--default_rr_set_group--ds_record--values--ds_key_algorithm) |
| `primary.default_rr_set_group.ds_record.values.key_tag` | [primary.default_rr_set_group.ds_record.values.key_tag](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values.md#schema-primary--default_rr_set_group--ds_record--values--key_tag) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values--sha1_digest.md#section) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest.digest](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values--sha1_digest.md#schema-primary--default_rr_set_group--ds_record--values--sha1_digest--digest) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values--sha256_digest.md#section) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest.digest](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values--sha256_digest.md#schema-primary--default_rr_set_group--ds_record--values--sha256_digest--digest) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values--sha384_digest.md#section) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest.digest](data-sources--dns_zone--properties--primary--default_rr_set_group--ds_record--values--sha384_digest.md#schema-primary--default_rr_set_group--ds_record--values--sha384_digest--digest) |
| `primary.default_rr_set_group.eui48_record` | [primary.default_rr_set_group.eui48_record](data-sources--dns_zone--properties--primary--default_rr_set_group--eui48_record.md#section) |
| `primary.default_rr_set_group.eui48_record.name` | [primary.default_rr_set_group.eui48_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--eui48_record.md#schema-primary--default_rr_set_group--eui48_record--name) |
| `primary.default_rr_set_group.eui48_record.value` | [primary.default_rr_set_group.eui48_record.value](data-sources--dns_zone--properties--primary--default_rr_set_group--eui48_record.md#schema-primary--default_rr_set_group--eui48_record--value) |
| `primary.default_rr_set_group.eui64_record` | [primary.default_rr_set_group.eui64_record](data-sources--dns_zone--properties--primary--default_rr_set_group--eui64_record.md#section) |
| `primary.default_rr_set_group.eui64_record.name` | [primary.default_rr_set_group.eui64_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--eui64_record.md#schema-primary--default_rr_set_group--eui64_record--name) |
| `primary.default_rr_set_group.eui64_record.value` | [primary.default_rr_set_group.eui64_record.value](data-sources--dns_zone--properties--primary--default_rr_set_group--eui64_record.md#schema-primary--default_rr_set_group--eui64_record--value) |
| `primary.default_rr_set_group.lb_record` | [primary.default_rr_set_group.lb_record](data-sources--dns_zone--properties--primary--default_rr_set_group--lb_record.md#section) |
| `primary.default_rr_set_group.lb_record.name` | [primary.default_rr_set_group.lb_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--lb_record.md#schema-primary--default_rr_set_group--lb_record--name) |
| `primary.default_rr_set_group.lb_record.value` | [primary.default_rr_set_group.lb_record.value](data-sources--dns_zone--properties--primary--default_rr_set_group--lb_record--value.md#section) |
| `primary.default_rr_set_group.lb_record.value.name` | [primary.default_rr_set_group.lb_record.value.name](data-sources--dns_zone--properties--primary--default_rr_set_group--lb_record--value.md#schema-primary--default_rr_set_group--lb_record--value--name) |
| `primary.default_rr_set_group.lb_record.value.namespace` | [primary.default_rr_set_group.lb_record.value.namespace](data-sources--dns_zone--properties--primary--default_rr_set_group--lb_record--value.md#schema-primary--default_rr_set_group--lb_record--value--namespace) |
| `primary.default_rr_set_group.lb_record.value.tenant` | [primary.default_rr_set_group.lb_record.value.tenant](data-sources--dns_zone--properties--primary--default_rr_set_group--lb_record--value.md#schema-primary--default_rr_set_group--lb_record--value--tenant) |
| `primary.default_rr_set_group.loc_record` | [primary.default_rr_set_group.loc_record](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record.md#section) |
| `primary.default_rr_set_group.loc_record.name` | [primary.default_rr_set_group.loc_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record.md#schema-primary--default_rr_set_group--loc_record--name) |
| `primary.default_rr_set_group.loc_record.values` | [primary.default_rr_set_group.loc_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#section) |
| `primary.default_rr_set_group.loc_record.values.altitude` | [primary.default_rr_set_group.loc_record.values.altitude](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--altitude) |
| `primary.default_rr_set_group.loc_record.values.horizontal_precision` | [primary.default_rr_set_group.loc_record.values.horizontal_precision](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--horizontal_precision) |
| `primary.default_rr_set_group.loc_record.values.latitude_degree` | [primary.default_rr_set_group.loc_record.values.latitude_degree](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--latitude_degree) |
| `primary.default_rr_set_group.loc_record.values.latitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.latitude_hemisphere](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--latitude_hemisphere) |
| `primary.default_rr_set_group.loc_record.values.latitude_minute` | [primary.default_rr_set_group.loc_record.values.latitude_minute](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--latitude_minute) |
| `primary.default_rr_set_group.loc_record.values.latitude_second` | [primary.default_rr_set_group.loc_record.values.latitude_second](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--latitude_second) |
| `primary.default_rr_set_group.loc_record.values.location_diameter` | [primary.default_rr_set_group.loc_record.values.location_diameter](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--location_diameter) |
| `primary.default_rr_set_group.loc_record.values.longitude_degree` | [primary.default_rr_set_group.loc_record.values.longitude_degree](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--longitude_degree) |
| `primary.default_rr_set_group.loc_record.values.longitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.longitude_hemisphere](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--longitude_hemisphere) |
| `primary.default_rr_set_group.loc_record.values.longitude_minute` | [primary.default_rr_set_group.loc_record.values.longitude_minute](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--longitude_minute) |
| `primary.default_rr_set_group.loc_record.values.longitude_second` | [primary.default_rr_set_group.loc_record.values.longitude_second](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--longitude_second) |
| `primary.default_rr_set_group.loc_record.values.vertical_precision` | [primary.default_rr_set_group.loc_record.values.vertical_precision](data-sources--dns_zone--properties--primary--default_rr_set_group--loc_record--values.md#schema-primary--default_rr_set_group--loc_record--values--vertical_precision) |
| `primary.default_rr_set_group.mx_record` | [primary.default_rr_set_group.mx_record](data-sources--dns_zone--properties--primary--default_rr_set_group--mx_record.md#section) |
| `primary.default_rr_set_group.mx_record.name` | [primary.default_rr_set_group.mx_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--mx_record.md#schema-primary--default_rr_set_group--mx_record--name) |
| `primary.default_rr_set_group.mx_record.values` | [primary.default_rr_set_group.mx_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--mx_record--values.md#section) |
| `primary.default_rr_set_group.mx_record.values.domain` | [primary.default_rr_set_group.mx_record.values.domain](data-sources--dns_zone--properties--primary--default_rr_set_group--mx_record--values.md#schema-primary--default_rr_set_group--mx_record--values--domain) |
| `primary.default_rr_set_group.mx_record.values.priority` | [primary.default_rr_set_group.mx_record.values.priority](data-sources--dns_zone--properties--primary--default_rr_set_group--mx_record--values.md#schema-primary--default_rr_set_group--mx_record--values--priority) |
| `primary.default_rr_set_group.naptr_record` | [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record.md#section) |
| `primary.default_rr_set_group.naptr_record.name` | [primary.default_rr_set_group.naptr_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record.md#schema-primary--default_rr_set_group--naptr_record--name) |
| `primary.default_rr_set_group.naptr_record.values` | [primary.default_rr_set_group.naptr_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record--values.md#section) |
| `primary.default_rr_set_group.naptr_record.values.flags` | [primary.default_rr_set_group.naptr_record.values.flags](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record--values.md#schema-primary--default_rr_set_group--naptr_record--values--flags) |
| `primary.default_rr_set_group.naptr_record.values.order` | [primary.default_rr_set_group.naptr_record.values.order](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record--values.md#schema-primary--default_rr_set_group--naptr_record--values--order) |
| `primary.default_rr_set_group.naptr_record.values.preference` | [primary.default_rr_set_group.naptr_record.values.preference](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record--values.md#schema-primary--default_rr_set_group--naptr_record--values--preference) |
| `primary.default_rr_set_group.naptr_record.values.regexp` | [primary.default_rr_set_group.naptr_record.values.regexp](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record--values.md#schema-primary--default_rr_set_group--naptr_record--values--regexp) |
| `primary.default_rr_set_group.naptr_record.values.replacement` | [primary.default_rr_set_group.naptr_record.values.replacement](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record--values.md#schema-primary--default_rr_set_group--naptr_record--values--replacement) |
| `primary.default_rr_set_group.naptr_record.values.service` | [primary.default_rr_set_group.naptr_record.values.service](data-sources--dns_zone--properties--primary--default_rr_set_group--naptr_record--values.md#schema-primary--default_rr_set_group--naptr_record--values--service) |
| `primary.default_rr_set_group.ns_record` | [primary.default_rr_set_group.ns_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ns_record.md#section) |
| `primary.default_rr_set_group.ns_record.name` | [primary.default_rr_set_group.ns_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--ns_record.md#schema-primary--default_rr_set_group--ns_record--name) |
| `primary.default_rr_set_group.ns_record.values` | [primary.default_rr_set_group.ns_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--ns_record.md#schema-primary--default_rr_set_group--ns_record--values) |
| `primary.default_rr_set_group.ptr_record` | [primary.default_rr_set_group.ptr_record](data-sources--dns_zone--properties--primary--default_rr_set_group--ptr_record.md#section) |
| `primary.default_rr_set_group.ptr_record.name` | [primary.default_rr_set_group.ptr_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--ptr_record.md#schema-primary--default_rr_set_group--ptr_record--name) |
| `primary.default_rr_set_group.ptr_record.values` | [primary.default_rr_set_group.ptr_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--ptr_record.md#schema-primary--default_rr_set_group--ptr_record--values) |
| `primary.default_rr_set_group.srv_record` | [primary.default_rr_set_group.srv_record](data-sources--dns_zone--properties--primary--default_rr_set_group--srv_record.md#section) |
| `primary.default_rr_set_group.srv_record.name` | [primary.default_rr_set_group.srv_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--srv_record.md#schema-primary--default_rr_set_group--srv_record--name) |
| `primary.default_rr_set_group.srv_record.values` | [primary.default_rr_set_group.srv_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--srv_record--values.md#section) |
| `primary.default_rr_set_group.srv_record.values.port` | [primary.default_rr_set_group.srv_record.values.port](data-sources--dns_zone--properties--primary--default_rr_set_group--srv_record--values.md#schema-primary--default_rr_set_group--srv_record--values--port) |
| `primary.default_rr_set_group.srv_record.values.priority` | [primary.default_rr_set_group.srv_record.values.priority](data-sources--dns_zone--properties--primary--default_rr_set_group--srv_record--values.md#schema-primary--default_rr_set_group--srv_record--values--priority) |
| `primary.default_rr_set_group.srv_record.values.target` | [primary.default_rr_set_group.srv_record.values.target](data-sources--dns_zone--properties--primary--default_rr_set_group--srv_record--values.md#schema-primary--default_rr_set_group--srv_record--values--target) |
| `primary.default_rr_set_group.srv_record.values.weight` | [primary.default_rr_set_group.srv_record.values.weight](data-sources--dns_zone--properties--primary--default_rr_set_group--srv_record--values.md#schema-primary--default_rr_set_group--srv_record--values--weight) |
| `primary.default_rr_set_group.sshfp_record` | [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record.md#section) |
| `primary.default_rr_set_group.sshfp_record.name` | [primary.default_rr_set_group.sshfp_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record.md#schema-primary--default_rr_set_group--sshfp_record--name) |
| `primary.default_rr_set_group.sshfp_record.values` | [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record--values.md#section) |
| `primary.default_rr_set_group.sshfp_record.values.algorithm` | [primary.default_rr_set_group.sshfp_record.values.algorithm](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record--values.md#schema-primary--default_rr_set_group--sshfp_record--values--algorithm) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record--values--sha1_fingerprint.md#section) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record--values--sha1_fingerprint.md#schema-primary--default_rr_set_group--sshfp_record--values--sha1_fingerprint--fingerprint) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record--values--sha256_fingerprint.md#section) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint](data-sources--dns_zone--properties--primary--default_rr_set_group--sshfp_record--values--sha256_fingerprint.md#schema-primary--default_rr_set_group--sshfp_record--values--sha256_fingerprint--fingerprint) |
| `primary.default_rr_set_group.tlsa_record` | [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--properties--primary--default_rr_set_group--tlsa_record.md#section) |
| `primary.default_rr_set_group.tlsa_record.name` | [primary.default_rr_set_group.tlsa_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--tlsa_record.md#schema-primary--default_rr_set_group--tlsa_record--name) |
| `primary.default_rr_set_group.tlsa_record.values` | [primary.default_rr_set_group.tlsa_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--tlsa_record--values.md#section) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_association_data` | [primary.default_rr_set_group.tlsa_record.values.certificate_association_data](data-sources--dns_zone--properties--primary--default_rr_set_group--tlsa_record--values.md#schema-primary--default_rr_set_group--tlsa_record--values--certificate_association_data) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_usage` | [primary.default_rr_set_group.tlsa_record.values.certificate_usage](data-sources--dns_zone--properties--primary--default_rr_set_group--tlsa_record--values.md#schema-primary--default_rr_set_group--tlsa_record--values--certificate_usage) |
| `primary.default_rr_set_group.tlsa_record.values.matching_type` | [primary.default_rr_set_group.tlsa_record.values.matching_type](data-sources--dns_zone--properties--primary--default_rr_set_group--tlsa_record--values.md#schema-primary--default_rr_set_group--tlsa_record--values--matching_type) |
| `primary.default_rr_set_group.tlsa_record.values.selector` | [primary.default_rr_set_group.tlsa_record.values.selector](data-sources--dns_zone--properties--primary--default_rr_set_group--tlsa_record--values.md#schema-primary--default_rr_set_group--tlsa_record--values--selector) |
| `primary.default_rr_set_group.ttl` | [primary.default_rr_set_group.ttl](data-sources--dns_zone--properties--primary--default_rr_set_group.md#schema-primary--default_rr_set_group--ttl) |
| `primary.default_rr_set_group.txt_record` | [primary.default_rr_set_group.txt_record](data-sources--dns_zone--properties--primary--default_rr_set_group--txt_record.md#section) |
| `primary.default_rr_set_group.txt_record.name` | [primary.default_rr_set_group.txt_record.name](data-sources--dns_zone--properties--primary--default_rr_set_group--txt_record.md#schema-primary--default_rr_set_group--txt_record--name) |
| `primary.default_rr_set_group.txt_record.values` | [primary.default_rr_set_group.txt_record.values](data-sources--dns_zone--properties--primary--default_rr_set_group--txt_record.md#schema-primary--default_rr_set_group--txt_record--values) |
| `primary.default_soa_parameters` | [primary.default_soa_parameters](data-sources--dns_zone--properties--primary--default_soa_parameters.md#section) |
| `primary.dnssec_mode` | [primary.dnssec_mode](data-sources--dns_zone--properties--primary--dnssec_mode.md#section) |
| `primary.dnssec_mode.disable_spec` | [primary.dnssec_mode.disable_spec](data-sources--dns_zone--properties--primary--dnssec_mode--disable_spec.md#section) |
| `primary.dnssec_mode.enable` | [primary.dnssec_mode.enable](data-sources--dns_zone--properties--primary--dnssec_mode--enable.md#section) |
| `primary.rr_set_group` | [primary.rr_set_group](data-sources--dns_zone--properties--primary--rr_set_group.md#section) |
| `primary.rr_set_group.metadata` | [primary.rr_set_group.metadata](data-sources--dns_zone--properties--primary--rr_set_group--metadata.md#section) |
| `primary.rr_set_group.metadata.description_spec` | [primary.rr_set_group.metadata.description_spec](data-sources--dns_zone--properties--primary--rr_set_group--metadata.md#schema-primary--rr_set_group--metadata--description_spec) |
| `primary.rr_set_group.metadata.name` | [primary.rr_set_group.metadata.name](data-sources--dns_zone--properties--primary--rr_set_group--metadata.md#schema-primary--rr_set_group--metadata--name) |
| `primary.rr_set_group.rr_set` | [primary.rr_set_group.rr_set](data-sources--dns_zone--properties--primary--rr_set_group--rr_set.md#section) |
| `primary.rr_set_group.rr_set.a_record` | [primary.rr_set_group.rr_set.a_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--a_record.md#section) |
| `primary.rr_set_group.rr_set.a_record.name` | [primary.rr_set_group.rr_set.a_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--a_record.md#schema-primary--rr_set_group--rr_set--a_record--name) |
| `primary.rr_set_group.rr_set.a_record.values` | [primary.rr_set_group.rr_set.a_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--a_record.md#schema-primary--rr_set_group--rr_set--a_record--values) |
| `primary.rr_set_group.rr_set.aaaa_record` | [primary.rr_set_group.rr_set.aaaa_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--aaaa_record.md#section) |
| `primary.rr_set_group.rr_set.aaaa_record.name` | [primary.rr_set_group.rr_set.aaaa_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--aaaa_record.md#schema-primary--rr_set_group--rr_set--aaaa_record--name) |
| `primary.rr_set_group.rr_set.aaaa_record.values` | [primary.rr_set_group.rr_set.aaaa_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--aaaa_record.md#schema-primary--rr_set_group--rr_set--aaaa_record--values) |
| `primary.rr_set_group.rr_set.afsdb_record` | [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record.md#section) |
| `primary.rr_set_group.rr_set.afsdb_record.name` | [primary.rr_set_group.rr_set.afsdb_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record.md#schema-primary--rr_set_group--rr_set--afsdb_record--name) |
| `primary.rr_set_group.rr_set.afsdb_record.values` | [primary.rr_set_group.rr_set.afsdb_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record--values.md#section) |
| `primary.rr_set_group.rr_set.afsdb_record.values.hostname` | [primary.rr_set_group.rr_set.afsdb_record.values.hostname](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record--values.md#schema-primary--rr_set_group--rr_set--afsdb_record--values--hostname) |
| `primary.rr_set_group.rr_set.afsdb_record.values.subtype` | [primary.rr_set_group.rr_set.afsdb_record.values.subtype](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--afsdb_record--values.md#schema-primary--rr_set_group--rr_set--afsdb_record--values--subtype) |
| `primary.rr_set_group.rr_set.alias_record` | [primary.rr_set_group.rr_set.alias_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--alias_record.md#section) |
| `primary.rr_set_group.rr_set.alias_record.value` | [primary.rr_set_group.rr_set.alias_record.value](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--alias_record.md#schema-primary--rr_set_group--rr_set--alias_record--value) |
| `primary.rr_set_group.rr_set.caa_record` | [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--caa_record.md#section) |
| `primary.rr_set_group.rr_set.caa_record.name` | [primary.rr_set_group.rr_set.caa_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--caa_record.md#schema-primary--rr_set_group--rr_set--caa_record--name) |
| `primary.rr_set_group.rr_set.caa_record.values` | [primary.rr_set_group.rr_set.caa_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--caa_record--values.md#section) |
| `primary.rr_set_group.rr_set.caa_record.values.flags` | [primary.rr_set_group.rr_set.caa_record.values.flags](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--caa_record--values.md#schema-primary--rr_set_group--rr_set--caa_record--values--flags) |
| `primary.rr_set_group.rr_set.caa_record.values.tag` | [primary.rr_set_group.rr_set.caa_record.values.tag](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--caa_record--values.md#schema-primary--rr_set_group--rr_set--caa_record--values--tag) |
| `primary.rr_set_group.rr_set.caa_record.values.value` | [primary.rr_set_group.rr_set.caa_record.values.value](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--caa_record--values.md#schema-primary--rr_set_group--rr_set--caa_record--values--value) |
| `primary.rr_set_group.rr_set.cds_record` | [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record.md#section) |
| `primary.rr_set_group.rr_set.cds_record.name` | [primary.rr_set_group.rr_set.cds_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record.md#schema-primary--rr_set_group--rr_set--cds_record--name) |
| `primary.rr_set_group.rr_set.cds_record.values` | [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record--values.md#section) |
| `primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record--values.md#schema-primary--rr_set_group--rr_set--cds_record--values--ds_key_algorithm) |
| `primary.rr_set_group.rr_set.cds_record.values.key_tag` | [primary.rr_set_group.rr_set.cds_record.values.key_tag](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record--values.md#schema-primary--rr_set_group--rr_set--cds_record--values--key_tag) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record--values--sha1_digest.md#section) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record--values--sha1_digest.md#schema-primary--rr_set_group--rr_set--cds_record--values--sha1_digest--digest) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record--values--sha256_digest.md#section) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record--values--sha256_digest.md#schema-primary--rr_set_group--rr_set--cds_record--values--sha256_digest--digest) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record--values--sha384_digest.md#section) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cds_record--values--sha384_digest.md#schema-primary--rr_set_group--rr_set--cds_record--values--sha384_digest--digest) |
| `primary.rr_set_group.rr_set.cert_record` | [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cert_record.md#section) |
| `primary.rr_set_group.rr_set.cert_record.name` | [primary.rr_set_group.rr_set.cert_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cert_record.md#schema-primary--rr_set_group--rr_set--cert_record--name) |
| `primary.rr_set_group.rr_set.cert_record.values` | [primary.rr_set_group.rr_set.cert_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cert_record--values.md#section) |
| `primary.rr_set_group.rr_set.cert_record.values.algorithm` | [primary.rr_set_group.rr_set.cert_record.values.algorithm](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cert_record--values.md#schema-primary--rr_set_group--rr_set--cert_record--values--algorithm) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_key_tag` | [primary.rr_set_group.rr_set.cert_record.values.cert_key_tag](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cert_record--values.md#schema-primary--rr_set_group--rr_set--cert_record--values--cert_key_tag) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_type` | [primary.rr_set_group.rr_set.cert_record.values.cert_type](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cert_record--values.md#schema-primary--rr_set_group--rr_set--cert_record--values--cert_type) |
| `primary.rr_set_group.rr_set.cert_record.values.certificate` | [primary.rr_set_group.rr_set.cert_record.values.certificate](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cert_record--values.md#schema-primary--rr_set_group--rr_set--cert_record--values--certificate) |
| `primary.rr_set_group.rr_set.cname_record` | [primary.rr_set_group.rr_set.cname_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cname_record.md#section) |
| `primary.rr_set_group.rr_set.cname_record.name` | [primary.rr_set_group.rr_set.cname_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cname_record.md#schema-primary--rr_set_group--rr_set--cname_record--name) |
| `primary.rr_set_group.rr_set.cname_record.value` | [primary.rr_set_group.rr_set.cname_record.value](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--cname_record.md#schema-primary--rr_set_group--rr_set--cname_record--value) |
| `primary.rr_set_group.rr_set.description_spec` | [primary.rr_set_group.rr_set.description_spec](data-sources--dns_zone--properties--primary--rr_set_group--rr_set.md#schema-primary--rr_set_group--rr_set--description_spec) |
| `primary.rr_set_group.rr_set.ds_record` | [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record.md#section) |
| `primary.rr_set_group.rr_set.ds_record.name` | [primary.rr_set_group.rr_set.ds_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record.md#schema-primary--rr_set_group--rr_set--ds_record--name) |
| `primary.rr_set_group.rr_set.ds_record.values` | [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values.md#section) |
| `primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values.md#schema-primary--rr_set_group--rr_set--ds_record--values--ds_key_algorithm) |
| `primary.rr_set_group.rr_set.ds_record.values.key_tag` | [primary.rr_set_group.rr_set.ds_record.values.key_tag](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values.md#schema-primary--rr_set_group--rr_set--ds_record--values--key_tag) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values--sha1_digest.md#section) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values--sha1_digest.md#schema-primary--rr_set_group--rr_set--ds_record--values--sha1_digest--digest) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values--sha256_digest.md#section) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values--sha256_digest.md#schema-primary--rr_set_group--rr_set--ds_record--values--sha256_digest--digest) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values--sha384_digest.md#section) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ds_record--values--sha384_digest.md#schema-primary--rr_set_group--rr_set--ds_record--values--sha384_digest--digest) |
| `primary.rr_set_group.rr_set.eui48_record` | [primary.rr_set_group.rr_set.eui48_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--eui48_record.md#section) |
| `primary.rr_set_group.rr_set.eui48_record.name` | [primary.rr_set_group.rr_set.eui48_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--eui48_record.md#schema-primary--rr_set_group--rr_set--eui48_record--name) |
| `primary.rr_set_group.rr_set.eui48_record.value` | [primary.rr_set_group.rr_set.eui48_record.value](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--eui48_record.md#schema-primary--rr_set_group--rr_set--eui48_record--value) |
| `primary.rr_set_group.rr_set.eui64_record` | [primary.rr_set_group.rr_set.eui64_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--eui64_record.md#section) |
| `primary.rr_set_group.rr_set.eui64_record.name` | [primary.rr_set_group.rr_set.eui64_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--eui64_record.md#schema-primary--rr_set_group--rr_set--eui64_record--name) |
| `primary.rr_set_group.rr_set.eui64_record.value` | [primary.rr_set_group.rr_set.eui64_record.value](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--eui64_record.md#schema-primary--rr_set_group--rr_set--eui64_record--value) |
| `primary.rr_set_group.rr_set.lb_record` | [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--lb_record.md#section) |
| `primary.rr_set_group.rr_set.lb_record.name` | [primary.rr_set_group.rr_set.lb_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--lb_record.md#schema-primary--rr_set_group--rr_set--lb_record--name) |
| `primary.rr_set_group.rr_set.lb_record.value` | [primary.rr_set_group.rr_set.lb_record.value](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--lb_record--value.md#section) |
| `primary.rr_set_group.rr_set.lb_record.value.name` | [primary.rr_set_group.rr_set.lb_record.value.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--lb_record--value.md#schema-primary--rr_set_group--rr_set--lb_record--value--name) |
| `primary.rr_set_group.rr_set.lb_record.value.namespace` | [primary.rr_set_group.rr_set.lb_record.value.namespace](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--lb_record--value.md#schema-primary--rr_set_group--rr_set--lb_record--value--namespace) |
| `primary.rr_set_group.rr_set.lb_record.value.tenant` | [primary.rr_set_group.rr_set.lb_record.value.tenant](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--lb_record--value.md#schema-primary--rr_set_group--rr_set--lb_record--value--tenant) |
| `primary.rr_set_group.rr_set.loc_record` | [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record.md#section) |
| `primary.rr_set_group.rr_set.loc_record.name` | [primary.rr_set_group.rr_set.loc_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record.md#schema-primary--rr_set_group--rr_set--loc_record--name) |
| `primary.rr_set_group.rr_set.loc_record.values` | [primary.rr_set_group.rr_set.loc_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#section) |
| `primary.rr_set_group.rr_set.loc_record.values.altitude` | [primary.rr_set_group.rr_set.loc_record.values.altitude](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--altitude) |
| `primary.rr_set_group.rr_set.loc_record.values.horizontal_precision` | [primary.rr_set_group.rr_set.loc_record.values.horizontal_precision](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--horizontal_precision) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.latitude_degree](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--latitude_degree) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--latitude_hemisphere) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.latitude_minute](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--latitude_minute) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_second` | [primary.rr_set_group.rr_set.loc_record.values.latitude_second](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--latitude_second) |
| `primary.rr_set_group.rr_set.loc_record.values.location_diameter` | [primary.rr_set_group.rr_set.loc_record.values.location_diameter](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--location_diameter) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.longitude_degree](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--longitude_degree) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--longitude_hemisphere) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.longitude_minute](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--longitude_minute) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_second` | [primary.rr_set_group.rr_set.loc_record.values.longitude_second](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--longitude_second) |
| `primary.rr_set_group.rr_set.loc_record.values.vertical_precision` | [primary.rr_set_group.rr_set.loc_record.values.vertical_precision](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--loc_record--values.md#schema-primary--rr_set_group--rr_set--loc_record--values--vertical_precision) |
| `primary.rr_set_group.rr_set.mx_record` | [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record.md#section) |
| `primary.rr_set_group.rr_set.mx_record.name` | [primary.rr_set_group.rr_set.mx_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record.md#schema-primary--rr_set_group--rr_set--mx_record--name) |
| `primary.rr_set_group.rr_set.mx_record.values` | [primary.rr_set_group.rr_set.mx_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record--values.md#section) |
| `primary.rr_set_group.rr_set.mx_record.values.domain` | [primary.rr_set_group.rr_set.mx_record.values.domain](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record--values.md#schema-primary--rr_set_group--rr_set--mx_record--values--domain) |
| `primary.rr_set_group.rr_set.mx_record.values.priority` | [primary.rr_set_group.rr_set.mx_record.values.priority](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--mx_record--values.md#schema-primary--rr_set_group--rr_set--mx_record--values--priority) |
| `primary.rr_set_group.rr_set.naptr_record` | [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record.md#section) |
| `primary.rr_set_group.rr_set.naptr_record.name` | [primary.rr_set_group.rr_set.naptr_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record.md#schema-primary--rr_set_group--rr_set--naptr_record--name) |
| `primary.rr_set_group.rr_set.naptr_record.values` | [primary.rr_set_group.rr_set.naptr_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record--values.md#section) |
| `primary.rr_set_group.rr_set.naptr_record.values.flags` | [primary.rr_set_group.rr_set.naptr_record.values.flags](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record--values.md#schema-primary--rr_set_group--rr_set--naptr_record--values--flags) |
| `primary.rr_set_group.rr_set.naptr_record.values.order` | [primary.rr_set_group.rr_set.naptr_record.values.order](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record--values.md#schema-primary--rr_set_group--rr_set--naptr_record--values--order) |
| `primary.rr_set_group.rr_set.naptr_record.values.preference` | [primary.rr_set_group.rr_set.naptr_record.values.preference](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record--values.md#schema-primary--rr_set_group--rr_set--naptr_record--values--preference) |
| `primary.rr_set_group.rr_set.naptr_record.values.regexp` | [primary.rr_set_group.rr_set.naptr_record.values.regexp](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record--values.md#schema-primary--rr_set_group--rr_set--naptr_record--values--regexp) |
| `primary.rr_set_group.rr_set.naptr_record.values.replacement` | [primary.rr_set_group.rr_set.naptr_record.values.replacement](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record--values.md#schema-primary--rr_set_group--rr_set--naptr_record--values--replacement) |
| `primary.rr_set_group.rr_set.naptr_record.values.service` | [primary.rr_set_group.rr_set.naptr_record.values.service](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--naptr_record--values.md#schema-primary--rr_set_group--rr_set--naptr_record--values--service) |
| `primary.rr_set_group.rr_set.ns_record` | [primary.rr_set_group.rr_set.ns_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ns_record.md#section) |
| `primary.rr_set_group.rr_set.ns_record.name` | [primary.rr_set_group.rr_set.ns_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ns_record.md#schema-primary--rr_set_group--rr_set--ns_record--name) |
| `primary.rr_set_group.rr_set.ns_record.values` | [primary.rr_set_group.rr_set.ns_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ns_record.md#schema-primary--rr_set_group--rr_set--ns_record--values) |
| `primary.rr_set_group.rr_set.ptr_record` | [primary.rr_set_group.rr_set.ptr_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ptr_record.md#section) |
| `primary.rr_set_group.rr_set.ptr_record.name` | [primary.rr_set_group.rr_set.ptr_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ptr_record.md#schema-primary--rr_set_group--rr_set--ptr_record--name) |
| `primary.rr_set_group.rr_set.ptr_record.values` | [primary.rr_set_group.rr_set.ptr_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--ptr_record.md#schema-primary--rr_set_group--rr_set--ptr_record--values) |
| `primary.rr_set_group.rr_set.srv_record` | [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--srv_record.md#section) |
| `primary.rr_set_group.rr_set.srv_record.name` | [primary.rr_set_group.rr_set.srv_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--srv_record.md#schema-primary--rr_set_group--rr_set--srv_record--name) |
| `primary.rr_set_group.rr_set.srv_record.values` | [primary.rr_set_group.rr_set.srv_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--srv_record--values.md#section) |
| `primary.rr_set_group.rr_set.srv_record.values.port` | [primary.rr_set_group.rr_set.srv_record.values.port](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--srv_record--values.md#schema-primary--rr_set_group--rr_set--srv_record--values--port) |
| `primary.rr_set_group.rr_set.srv_record.values.priority` | [primary.rr_set_group.rr_set.srv_record.values.priority](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--srv_record--values.md#schema-primary--rr_set_group--rr_set--srv_record--values--priority) |
| `primary.rr_set_group.rr_set.srv_record.values.target` | [primary.rr_set_group.rr_set.srv_record.values.target](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--srv_record--values.md#schema-primary--rr_set_group--rr_set--srv_record--values--target) |
| `primary.rr_set_group.rr_set.srv_record.values.weight` | [primary.rr_set_group.rr_set.srv_record.values.weight](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--srv_record--values.md#schema-primary--rr_set_group--rr_set--srv_record--values--weight) |
| `primary.rr_set_group.rr_set.sshfp_record` | [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record.md#section) |
| `primary.rr_set_group.rr_set.sshfp_record.name` | [primary.rr_set_group.rr_set.sshfp_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record.md#schema-primary--rr_set_group--rr_set--sshfp_record--name) |
| `primary.rr_set_group.rr_set.sshfp_record.values` | [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record--values.md#section) |
| `primary.rr_set_group.rr_set.sshfp_record.values.algorithm` | [primary.rr_set_group.rr_set.sshfp_record.values.algorithm](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record--values.md#schema-primary--rr_set_group--rr_set--sshfp_record--values--algorithm) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record--values--sha1_fingerprint.md#section) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record--values--sha1_fingerprint.md#schema-primary--rr_set_group--rr_set--sshfp_record--values--sha1_fingerprint--fingerprint) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record--values--sha256_fingerprint.md#section) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--sshfp_record--values--sha256_fingerprint.md#schema-primary--rr_set_group--rr_set--sshfp_record--values--sha256_fingerprint--fingerprint) |
| `primary.rr_set_group.rr_set.tlsa_record` | [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--tlsa_record.md#section) |
| `primary.rr_set_group.rr_set.tlsa_record.name` | [primary.rr_set_group.rr_set.tlsa_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--tlsa_record.md#schema-primary--rr_set_group--rr_set--tlsa_record--name) |
| `primary.rr_set_group.rr_set.tlsa_record.values` | [primary.rr_set_group.rr_set.tlsa_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--tlsa_record--values.md#section) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--tlsa_record--values.md#schema-primary--rr_set_group--rr_set--tlsa_record--values--certificate_association_data) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--tlsa_record--values.md#schema-primary--rr_set_group--rr_set--tlsa_record--values--certificate_usage) |
| `primary.rr_set_group.rr_set.tlsa_record.values.matching_type` | [primary.rr_set_group.rr_set.tlsa_record.values.matching_type](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--tlsa_record--values.md#schema-primary--rr_set_group--rr_set--tlsa_record--values--matching_type) |
| `primary.rr_set_group.rr_set.tlsa_record.values.selector` | [primary.rr_set_group.rr_set.tlsa_record.values.selector](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--tlsa_record--values.md#schema-primary--rr_set_group--rr_set--tlsa_record--values--selector) |
| `primary.rr_set_group.rr_set.ttl` | [primary.rr_set_group.rr_set.ttl](data-sources--dns_zone--properties--primary--rr_set_group--rr_set.md#schema-primary--rr_set_group--rr_set--ttl) |
| `primary.rr_set_group.rr_set.txt_record` | [primary.rr_set_group.rr_set.txt_record](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--txt_record.md#section) |
| `primary.rr_set_group.rr_set.txt_record.name` | [primary.rr_set_group.rr_set.txt_record.name](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--txt_record.md#schema-primary--rr_set_group--rr_set--txt_record--name) |
| `primary.rr_set_group.rr_set.txt_record.values` | [primary.rr_set_group.rr_set.txt_record.values](data-sources--dns_zone--properties--primary--rr_set_group--rr_set--txt_record.md#schema-primary--rr_set_group--rr_set--txt_record--values) |
| `primary.soa_parameters` | [primary.soa_parameters](data-sources--dns_zone--properties--primary--soa_parameters.md#section) |
| `primary.soa_parameters.expire` | [primary.soa_parameters.expire](data-sources--dns_zone--properties--primary--soa_parameters.md#schema-primary--soa_parameters--expire) |
| `primary.soa_parameters.negative_ttl` | [primary.soa_parameters.negative_ttl](data-sources--dns_zone--properties--primary--soa_parameters.md#schema-primary--soa_parameters--negative_ttl) |
| `primary.soa_parameters.refresh` | [primary.soa_parameters.refresh](data-sources--dns_zone--properties--primary--soa_parameters.md#schema-primary--soa_parameters--refresh) |
| `primary.soa_parameters.retry` | [primary.soa_parameters.retry](data-sources--dns_zone--properties--primary--soa_parameters.md#schema-primary--soa_parameters--retry) |
| `primary.soa_parameters.ttl` | [primary.soa_parameters.ttl](data-sources--dns_zone--properties--primary--soa_parameters.md#schema-primary--soa_parameters--ttl) |
| `secondary` | [secondary](data-sources--dns_zone--properties--secondary.md#section) |
| `secondary.primary_servers` | [secondary.primary_servers](data-sources--dns_zone--properties--secondary.md#schema-secondary--primary_servers) |
| `secondary.tsig_key_algorithm` | [secondary.tsig_key_algorithm](data-sources--dns_zone--properties--secondary.md#schema-secondary--tsig_key_algorithm) |
| `secondary.tsig_key_name` | [secondary.tsig_key_name](data-sources--dns_zone--properties--secondary.md#schema-secondary--tsig_key_name) |
| `secondary.tsig_key_value` | [secondary.tsig_key_value](data-sources--dns_zone--properties--secondary--tsig_key_value.md#section) |
| `secondary.tsig_key_value.blindfold_secret_info` | [secondary.tsig_key_value.blindfold_secret_info](data-sources--dns_zone--properties--secondary--tsig_key_value--blindfold_secret_info.md#section) |
| `secondary.tsig_key_value.blindfold_secret_info.decryption_provider` | [secondary.tsig_key_value.blindfold_secret_info.decryption_provider](data-sources--dns_zone--properties--secondary--tsig_key_value--blindfold_secret_info.md#schema-secondary--tsig_key_value--blindfold_secret_info--decryption_provider) |
| `secondary.tsig_key_value.blindfold_secret_info.location` | [secondary.tsig_key_value.blindfold_secret_info.location](data-sources--dns_zone--properties--secondary--tsig_key_value--blindfold_secret_info.md#schema-secondary--tsig_key_value--blindfold_secret_info--location) |
| `secondary.tsig_key_value.blindfold_secret_info.store_provider` | [secondary.tsig_key_value.blindfold_secret_info.store_provider](data-sources--dns_zone--properties--secondary--tsig_key_value--blindfold_secret_info.md#schema-secondary--tsig_key_value--blindfold_secret_info--store_provider) |
| `secondary.tsig_key_value.clear_secret_info` | [secondary.tsig_key_value.clear_secret_info](data-sources--dns_zone--properties--secondary--tsig_key_value--clear_secret_info.md#section) |
| `secondary.tsig_key_value.clear_secret_info.provider_ref` | [secondary.tsig_key_value.clear_secret_info.provider_ref](data-sources--dns_zone--properties--secondary--tsig_key_value--clear_secret_info.md#schema-secondary--tsig_key_value--clear_secret_info--provider_ref) |
| `secondary.tsig_key_value.clear_secret_info.url` | [secondary.tsig_key_value.clear_secret_info.url](data-sources--dns_zone--properties--secondary--tsig_key_value--clear_secret_info.md#schema-secondary--tsig_key_value--clear_secret_info--url) |

## Next pages

- [primary](data-sources--dns_zone--properties--primary.md)
- [secondary](data-sources--dns_zone--properties--secondary.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
