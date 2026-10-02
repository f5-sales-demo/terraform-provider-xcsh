---
page_title: "items.get_spec.infra.hw_info.cpu"
subcategory: ""
description: "CPU Information. CPU information."
xcsh_docs: {"aliases": ["items get spec infra hw info cpu"], "body_bytes": 2409, "body_sha256": "sha256:938e797e455918943181a21c368b1e9f89a6ed3400bb9d0534e9772dcc42f36e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info", "path": "documentation/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/cpu/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2003310202210232-0330113122010301-3003022101231220-2201323033230211-2011031321102300-1320113223312201-3033031133032130-0220331300301113", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "hw_info", "cpu"], "schema_version": 1, "sections": [{"aliases": ["cache"], "anchor": "schema-items--get_spec--infra--hw_info--cpu--cache", "description": "Cache. CPU cache size in KB.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "cpu", "cache"], "syntax": "attribute", "type": "number"}, {"aliases": ["cores"], "anchor": "schema-items--get_spec--infra--hw_info--cpu--cores", "description": "Cores. Number of physical CPU cores.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "cpu", "cores"], "syntax": "attribute", "type": "number"}, {"aliases": ["cpus"], "anchor": "schema-items--get_spec--infra--hw_info--cpu--cpus", "description": "CPUs. Number of physical CPUs.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "cpu", "cpus"], "syntax": "attribute", "type": "number"}, {"aliases": ["model"], "anchor": "schema-items--get_spec--infra--hw_info--cpu--model", "description": "Model. CPU model", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "cpu", "model"], "syntax": "attribute", "type": "string"}, {"aliases": ["speed"], "anchor": "schema-items--get_spec--infra--hw_info--cpu--speed", "description": "Speed. CPU clock rate in MHz.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "cpu", "speed"], "syntax": "attribute", "type": "number"}, {"aliases": ["threads"], "anchor": "schema-items--get_spec--infra--hw_info--cpu--threads", "description": "Threads. Number of logical (HT) CPU cores.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "cpu", "threads"], "syntax": "attribute", "type": "number"}, {"aliases": ["vendor"], "anchor": "schema-items--get_spec--infra--hw_info--cpu--vendor", "description": "Vendor. CPU vendor.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:get_spec:infra:hw_info:cpu", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "hw_info", "cpu", "vendor"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/cpu/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "CPU Information. CPU information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.hw_info.cpu

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/)
- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- items.get_spec.infra.hw_info.cpu

<a id="section"></a>

Type: `"single"`. Computed.

CPU Information. CPU information.

## Direct properties

<a id="schema-items--get_spec--infra--hw_info--cpu--cache"></a>

### cache property

Type: `"number"`. Computed.

Cache. CPU cache size in KB.

<a id="schema-items--get_spec--infra--hw_info--cpu--cores"></a>

### cores property

Type: `"number"`. Computed.

Cores. Number of physical CPU cores.

<a id="schema-items--get_spec--infra--hw_info--cpu--cpus"></a>

### cpus property

Type: `"number"`. Computed.

CPUs. Number of physical CPUs.

<a id="schema-items--get_spec--infra--hw_info--cpu--model"></a>

### model property

Type: `"string"`. Computed.

Model. CPU model

<a id="schema-items--get_spec--infra--hw_info--cpu--speed"></a>

### speed property

Type: `"number"`. Computed.

Speed. CPU clock rate in MHz.

<a id="schema-items--get_spec--infra--hw_info--cpu--threads"></a>

### threads property

Type: `"number"`. Computed.

Threads. Number of logical (HT) CPU cores.

<a id="schema-items--get_spec--infra--hw_info--cpu--vendor"></a>

### vendor property

Type: `"string"`. Computed.

Vendor. CPU vendor.

## Next pages

- [items.get_spec.infra.hw_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/get_spec/infra/hw_info/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
