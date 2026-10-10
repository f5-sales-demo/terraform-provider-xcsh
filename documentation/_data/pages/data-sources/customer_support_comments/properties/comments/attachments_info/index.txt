---
page_title: "comments.attachments_info"
subcategory: ""
description: "Information about any attachments (such as screenshots, plain text files) the comment can have."
xcsh_docs: {"aliases": ["comments attachments info"], "body_bytes": 1490, "body_sha256": "sha256:df67df7c8240e37e515d3d13b566504ddd44f9fbe966486eb484507647c90f57", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "parent_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "path": "documentation/data-sources/customer_support_comments/properties/comments/attachments_info/index.md", "product": "distributed-cloud", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1123201102321012-0010233330122000-2321201222222020-2123030310113230-2200012331021303-2311011102210123-3300232002210220-0110001102200201", "registry_path": "docs/guides/data-sources--customer_support_comments--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["comments", "attachments_info"], "schema_version": 1, "sections": [{"aliases": ["comments attachments info attachment"], "anchor": "schema-comments--attachments_info--attachment", "description": "Any binary attachment (such as screenshots, plain text files, PDFs) encoded as base64 if used over HTTP.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachments_info", "attachment"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments attachments info content type"], "anchor": "schema-comments--attachments_info--content_type", "description": "MIME content type of the attachment. Helps the UI to properly display the data.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachments_info", "content_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments attachments info filename"], "anchor": "schema-comments--attachments_info--filename", "description": "Filename of the attachment as provided by the caller.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachments_info", "filename"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments attachments info tp id"], "anchor": "schema-comments--attachments_info--tp_id", "description": "Optional ID as assigned by the third-party actually storing the data.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachments_info", "tp_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/properties/comments/attachments_info/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Information about any attachments (such as screenshots, plain text files) the comment can have.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# comments.attachments_info

Breadcrumbs:

- [xcsh_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/)
- [comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/)
- comments.attachments_info

<a id="section"></a>

Type: `"list"`. Computed.

Information about any attachments (such as screenshots, plain text files) the comment can have.

## Direct properties

<a id="schema-comments--attachments_info--attachment"></a>

### attachment property

Type: `"string"`. Computed.

Any binary attachment (such as screenshots, plain text files, PDFs) encoded as base64 if used over
HTTP.

<a id="schema-comments--attachments_info--content_type"></a>

### content_type property

Type: `"string"`. Computed.

MIME content type of the attachment. Helps the UI to properly display the data.

<a id="schema-comments--attachments_info--filename"></a>

### filename property

Type: `"string"`. Computed.

Filename of the attachment as provided by the caller.

<a id="schema-comments--attachments_info--tp_id"></a>

### tp_id property

Type: `"string"`. Computed.

Optional ID as assigned by the third-party actually storing the data.
