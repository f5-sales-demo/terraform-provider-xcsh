---
page_title: "comments"
subcategory: ""
description: "List of comments on the customer support ticket."
xcsh_docs: {"aliases": ["comments"], "body_bytes": 1980, "body_sha256": "sha256:d60d2c6e6f0d47c0018bb4fb7dd4dcbd20d52fe47852bd4815fb13166e91ea22", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "parent_id": "xcsh-docs:data-sources:customer_support_comments:reference", "path": "documentation/data-sources/customer_support_comments/properties/comments/index.md", "product": "distributed-cloud", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0212003003321222-1012231332100122-2323101111111201-2022033202330130-0133312221320120-2130233133303021-1123002210222130-1233020223320121", "registry_path": "docs/guides/data-sources--customer_support_comments--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["comments"], "schema_version": 1, "sections": [{"aliases": ["comments attachment ids"], "anchor": "schema-comments--attachment_ids", "description": "Third party ID of any attachment related to this ticket comment.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "attachment_ids"], "syntax": "attribute", "type": "list"}, {"aliases": ["comments attachments info"], "anchor": "section", "description": "Information about any attachments (such as screenshots, plain text files) the comment can have.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments:attachments_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["comments", "attachments_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["comments author email"], "anchor": "schema-comments--author_email", "description": "Email. Email of the author of the comment.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "author_email"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments author name"], "anchor": "schema-comments--author_name", "description": "Author. Author of the comment (as a name)", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "author_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments comment id"], "anchor": "schema-comments--comment_id", "description": "ID assigned to this comment by support provider.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "comment_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments created at"], "anchor": "schema-comments--created_at", "description": "At. Comment creation time.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "created_at"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments html"], "anchor": "schema-comments--html", "description": "Comment. Comment body as HTML.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "html"], "syntax": "attribute", "type": "string"}, {"aliases": ["comments plain text"], "anchor": "schema-comments--plain_text", "description": "Comment. Comment body as plain text.", "document_id": "xcsh-docs:data-sources:customer_support_comments:properties:comments", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["comments", "plain_text"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/properties/comments/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "List of comments on the customer support ticket.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# comments

Breadcrumbs:

- [xcsh_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/)
- comments

<a id="section"></a>

Type: `"list"`. Computed.

List of comments on the customer support ticket.

## Direct properties

<a id="schema-comments--attachment_ids"></a>

### attachment_ids property

Type: `["list", "string"]`. Computed.

Third party ID of any attachment related to this ticket comment.

- [attachments_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/comments/attachments_info/): complete subsection reference.

<a id="schema-comments--author_email"></a>

### author_email property

Type: `"string"`. Computed.

Email. Email of the author of the comment.

<a id="schema-comments--author_name"></a>

### author_name property

Type: `"string"`. Computed.

Author. Author of the comment (as a name)

<a id="schema-comments--comment_id"></a>

### comment_id property

Type: `"string"`. Computed.

ID assigned to this comment by support provider.

<a id="schema-comments--created_at"></a>

### created_at property

Type: `"string"`. Computed.

At. Comment creation time.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d{3})?Z?$`),
    ""),
}
```

<a id="schema-comments--html"></a>

### html property

Type: `"string"`. Computed.

Comment. Comment body as HTML.

<a id="schema-comments--plain_text"></a>

### plain_text property

Type: `"string"`. Computed.

Comment. Comment body as plain text.
