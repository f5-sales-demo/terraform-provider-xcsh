---
page_title: "xcsh_dns_zone_cryptokeys"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["dns zone cryptokeys"], "body_bytes": 1213, "body_sha256": "sha256:b26cf803f1868029962b5cccafbfcba025f261d9f105ccf40f1d5732bc949d2c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:dns_zone_cryptokeys:reference", "xcsh-docs:data-sources:dns_zone_cryptokeys:examples"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_zone_cryptokeys:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone_cryptokeys:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/dns_zone_cryptokeys/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_cryptokeys", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3311023113031202-1010300221212322-1133102011021320-0033123312013310-0302320000020011-0000120213001132-1111132220110232-1220301011030302", "registry_path": "docs/data-sources/dns_zone_cryptokeys.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone_cryptokeys/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_zone_cryptokeys

Breadcrumbs:

- xcsh_dns_zone_cryptokeys

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZoneCryptokeys DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_dns_zone_cryptokeys" "example" {
}

output "dns_zone_cryptokeys_result" {
  value = data.xcsh_dns_zone_cryptokeys.example
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone_cryptokeys/examples/)
