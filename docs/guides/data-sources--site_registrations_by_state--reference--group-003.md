---
page_title: "xcsh_site_registrations_by_state reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_state reference."
---

# xcsh_site_registrations_by_state reference

<a id="canonical-9336ac9e42b4db6f44e5bbdc8a18403b706168e88cd264cfbcfb25404abf003a"></a>

## Direct properties — items.object.system_metadata.initializers / c4c77f985244 / 3

- [pending](data-sources--site_registrations_by_state--reference--group-003.md#canonical-f16c7eb3ff02420763a206485419d532e0b1cfc1638b90cb42cac5c3cbd6ee2d): complete subsection reference.

- [result](data-sources--site_registrations_by_state--reference--group-003.md#canonical-afac902a61501586c49bc4d808cee7e72b7055d46dd4c1c120216283f6629270): complete subsection reference.

<a id="canonical-f5299604aefd4c2a04d641b9d4a22d317a3283be38de1750b6942262bfba67ab"></a>

## Next pages — items.object.system_metadata.initializers / c4c77f985244 / 4

- [items.object.system_metadata.initializers.pending](data-sources--site_registrations_by_state--reference--group-003.md#canonical-f16c7eb3ff02420763a206485419d532e0b1cfc1638b90cb42cac5c3cbd6ee2d)
- [items.object.system_metadata.initializers.result](data-sources--site_registrations_by_state--reference--group-003.md#canonical-afac902a61501586c49bc4d808cee7e72b7055d46dd4c1c120216283f6629270)
- [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-877448804fbc002b6806788021a0e670d7179f38b80a98f00469b87c1f686f82)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-f16c7eb3ff02420763a206485419d532e0b1cfc1638b90cb42cac5c3cbd6ee2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eeb7fa73d6c6f0c4f50076c6b73725f4ef7f6b34a6179447ecdb78b53f02255d"></a>

## items.object.system_metadata.initializers.pending — items.object.system_metadata.initializers.pending / 5b334940f779 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6b2a79c49bdb8022e6123c3c9b65430899eba4c3cdce0da3fb5aa24a03144383)
- [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-877448804fbc002b6806788021a0e670d7179f38b80a98f00469b87c1f686f82)
- [items.object.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-002.md#canonical-99e4872bf9a3e543cf18bbafe0c128c8a8f922cf6fe695715029b95e44a2bdee)
- items.object.system_metadata.initializers.pending

<a id="canonical-e7269a7c93daf6350ddf4108408227058332a74e017a5ce164bb65ef2eba8605"></a>

Type: `"list"`. Computed.

Pending is a list of initializers that must execute in order before this object is initialized. When
the last pending initializer is removed, and no failing result is set, the initializers struct will
be set to nil and the object is considered as initialized and visible to all clients.

<a id="canonical-d80288353076ab06bb181d08e12c2ec8b2cfbe2915cc386c87211d36b9a46bce"></a>

## Direct properties — items.object.system_metadata.initializers.pending / 5b334940f779 / 3

<a id="canonical-27fcb5f779a437fd6d83398fb9a9634e16e97c27054e07cd76ad8497828c8cbe"></a>

<a id="canonical-f2cd2169812bf5cfa0727992411b82cfd39924307b8ffa984355c0879b184280"></a>

## name property — items.object.system_metadata.initializers.pending / 5b334940f779 / 4

Type: `"string"`. Computed.

Name of the service that is responsible for initializing this object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-f31b77c6d21e2362abbf4b6e14d7b2d2fc190772c61470162e415ddacdf3e605"></a>

## Next pages — items.object.system_metadata.initializers.pending / 5b334940f779 / 5

- [items.object.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-002.md#canonical-99e4872bf9a3e543cf18bbafe0c128c8a8f922cf6fe695715029b95e44a2bdee)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-afac902a61501586c49bc4d808cee7e72b7055d46dd4c1c120216283f6629270"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88dd90d0f4b14de0abb4cc40659cc448e17fc1a8ef2e52a321f2665220d95c91"></a>

## items.object.system_metadata.initializers.result — items.object.system_metadata.initializers.result / f7aaf0c19374 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6b2a79c49bdb8022e6123c3c9b65430899eba4c3cdce0da3fb5aa24a03144383)
- [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-877448804fbc002b6806788021a0e670d7179f38b80a98f00469b87c1f686f82)
- [items.object.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-002.md#canonical-99e4872bf9a3e543cf18bbafe0c128c8a8f922cf6fe695715029b95e44a2bdee)
- items.object.system_metadata.initializers.result

<a id="canonical-40b2d22aa5545bc7edb167d58ffd9f1773455f58318e3f9e825aa26499239f4d"></a>

Type: `"single"`. Computed.

Status is a return value for calls that don't return other objects.

<a id="canonical-fb7a5d96159e73c9f1240e6b39437d79ad132b8aacf34cbcb96054c691bca5a9"></a>

## Direct properties — items.object.system_metadata.initializers.result / f7aaf0c19374 / 3

<a id="canonical-f561fad416bb3f7e0bbf3b7415e0b821712b1a265534a58069944bf07194c2c1"></a>

<a id="canonical-68c35810d0ee563d3a6432e7cffa36b7cdf4d14ce6d57c780e614a11baec49ab"></a>

## code property — items.object.system_metadata.initializers.result / f7aaf0c19374 / 4

Type: `"number"`. Computed.

Suggested HTTP return code for this status, 0 if not set.

<a id="canonical-7fe0b39ab84b5f41362a4c9013d1af4113ed15d5c46ea9c48c116946afead738"></a>

<a id="canonical-00fc381ac046f3a764147ce2cd185b416c7553b552c1e294f13518b7b6cf53e5"></a>

## reason property — items.object.system_metadata.initializers.result / f7aaf0c19374 / 5

Type: `"string"`. Computed.

Human-readable description of why this operation is in the 'Failure' status. If this value is empty
there is no information available.

<a id="canonical-6d10e060e21ab13bd50d2956f3aa2fb8d5b084e30025cc3829a8340d01dc0937"></a>

<a id="canonical-789424c1708bdc95d2af5f5f2191f378b6d9a4403570052d1eace4f61a70b4a8"></a>

## status property — items.object.system_metadata.initializers.result / f7aaf0c19374 / 6

Type: `"string"`. Computed.

Status of the operation. One of: 'Success' or 'Failure'.

<a id="canonical-27578864285f5cd9d206a88dfdb5d3503cad5cabbabc60fee8b6b27f4ceea378"></a>

## Next pages — items.object.system_metadata.initializers.result / f7aaf0c19374 / 7

- [items.object.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-002.md#canonical-99e4872bf9a3e543cf18bbafe0c128c8a8f922cf6fe695715029b95e44a2bdee)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-2ee7885c16a6c41fb66e5359a844283f28fbab4e916814336c8adb4b33ce4084"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f1d15aeed53cc8076dafeb25a6c9c6245f9dac9ac26b02f66859a1b57264e1a"></a>

## items.object.system_metadata.labels — items.object.system_metadata.labels / f253febb7d3b / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6b2a79c49bdb8022e6123c3c9b65430899eba4c3cdce0da3fb5aa24a03144383)
- [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-877448804fbc002b6806788021a0e670d7179f38b80a98f00469b87c1f686f82)
- items.object.system_metadata.labels

<a id="canonical-0bcb9d25c56f96b7de6a207a0a5817c61006debe6c7adc4e51c780e54c3f57c2"></a>

Type: `"single"`. Computed.

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the operator or software. Values here can be interpreted by software(backend or
frontend) to enable certain behavior e.g. Things marked as soft-deleted(restorable).

<a id="canonical-096d88b6d0d1989a5af5f751ae45ab22c234d796356f39b09796609c3522ad15"></a>

## Direct properties — items.object.system_metadata.labels / f253febb7d3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a345c17b73b2a0a70dcd326b19a9075f2c9eff8a82117dd998b580e67199a508"></a>

## Next pages — items.object.system_metadata.labels / f253febb7d3b / 4

- [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-877448804fbc002b6806788021a0e670d7179f38b80a98f00469b87c1f686f82)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-2aa3fa530d64b62c5fe9889177649a9b1ca81cf213d61fd992a542233211ce68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0770d69f2699a0afb6cf0fabc414859786903d3c9b4f416505899cde04e314b0"></a>

## items.object.system_metadata.namespace — items.object.system_metadata.namespace / cddc9bf644d3 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6b2a79c49bdb8022e6123c3c9b65430899eba4c3cdce0da3fb5aa24a03144383)
- [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-877448804fbc002b6806788021a0e670d7179f38b80a98f00469b87c1f686f82)
- items.object.system_metadata.namespace

<a id="canonical-06aaa240001d9b8379715572ea969b79d8ef416be24077e1b789890b664889c6"></a>

Type: `"list"`. Computed.

The namespace this object belongs to. This is populated by the service based on the
metadata.namespace field when an object is created.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

<a id="canonical-4ee2f6222b99a7b806c472e3347cf87c23b3258f9492617d76bc0ea432fdbdd8"></a>

## Direct properties — items.object.system_metadata.namespace / cddc9bf644d3 / 3

<a id="canonical-4e6493cc1f999f9c20a7436290e195742172a0aac40dc163743bf4283b55c099"></a>

<a id="canonical-a1374b0a92abf3e1721de8a25c563c9cc0bb9be01818ef486cca7d57254582e1"></a>

## kind property — items.object.system_metadata.namespace / cddc9bf644d3 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="canonical-c90918569e14a85a2afdce69d9982ae3859ff5d6e4a579c7ac47758d7830b365"></a>

<a id="canonical-175136da837760d18862ebb6f3c72f06ef6479bc95c541ce63e8802ba5af40c6"></a>

## name property — items.object.system_metadata.namespace / cddc9bf644d3 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="canonical-f1c9a4a07d90696b390fa15b034ea52d29c935845118b25f107ebb1d38c42050"></a>

<a id="canonical-8dc11dd237d73d365340860ab6a3c84c8c2710845c5ad56e080ffd97bffc2903"></a>

## namespace property — items.object.system_metadata.namespace / cddc9bf644d3 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-38a2226f176f16aab130ce69c0a8686cbddf8bb01411898db197434d2e04b9ab"></a>

<a id="canonical-e1b67f211f5e7c18cae5bd25ec3b4bf6ef4cd160c8b125652b14219ba6a9b909"></a>

## tenant property — items.object.system_metadata.namespace / cddc9bf644d3 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="canonical-19348bbc2becbb6442f33372ba7b2ad95c4e8158eaeb2c81e35b634fe01177ea"></a>

<a id="canonical-5e2a30e1ac24992c3c3011eb4475ff40a0cdd8b47b29952c6816b36ec2b28de7"></a>

## uid property — items.object.system_metadata.namespace / cddc9bf644d3 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

<a id="canonical-5fb6ed9af92fd724ad4a24f533a124384e708acad8b6d80f911f046f9f9e40c2"></a>

## Next pages — items.object.system_metadata.namespace / cddc9bf644d3 / 9

- [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-877448804fbc002b6806788021a0e670d7179f38b80a98f00469b87c1f686f82)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-7e6e17cd791887ed4d0db55c10f1688ab41c2dabce00a3e7c05bfea6906beeb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3be84a9fe14753b69087a6da1ad29100772b25fb7f96067c926b97dccfed190d"></a>

## items.object.system_metadata.owner_view — items.object.system_metadata.owner_view / ca96e2798127 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.object](data-sources--site_registrations_by_state--reference--group-002.md#canonical-6b2a79c49bdb8022e6123c3c9b65430899eba4c3cdce0da3fb5aa24a03144383)
- [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-877448804fbc002b6806788021a0e670d7179f38b80a98f00469b87c1f686f82)
- items.object.system_metadata.owner_view

<a id="canonical-a0e7b40d37dd980b0e7c1fd21b254a6d4a15044fd1f819e396264861b0ccaa91"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

<a id="canonical-2790ab469219d1fb1bf80c24711765c493c70d62856383ddd2e901c60ab5265a"></a>

## Direct properties — items.object.system_metadata.owner_view / ca96e2798127 / 3

<a id="canonical-8589f28c8836558e2b2bc3302ae1bdd12099492a4af5ba121a3cf418702c45c3"></a>

<a id="canonical-a628ca405cc6ca27d0f1f3147cdf07ac5ba1f9f292a2ff2373db0a705c8485fe"></a>

## kind property — items.object.system_metadata.owner_view / ca96e2798127 / 4

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="canonical-8b400ac879b5509b90059b01b44ca1b55fcbb83942de9421b73734bfbf0bf5eb"></a>

<a id="canonical-e8dc1f8144df08e586b35076b75d9ae7b6886971aea0370e610e908291e686b9"></a>

## name property — items.object.system_metadata.owner_view / ca96e2798127 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-1120ea5704a8d5f08944dce88ea999f2018d7005a2d8a14373bac735a154fba7"></a>

<a id="canonical-6e400f12840a142727a9ef0415e2f63ed39fa0f7669834f23c5e2c9d2b24edc2"></a>

## namespace property — items.object.system_metadata.owner_view / ca96e2798127 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-44db5d9b0f873fb458d30b1529d7969f9e512b6f793b333d40b7b50ce0cd4b72"></a>

<a id="canonical-fcc6c5bbd5ffa3646c105819ae51b700507d6a7bd6294ef8ac6f1d1da1dab395"></a>

## uid property — items.object.system_metadata.owner_view / ca96e2798127 / 7

Type: `"string"`. Computed.

UID. UID of the view object.

<a id="canonical-8b8a1d87a572353979b100fa17affb2677fd709e5fa1a66ce761afb0c82c976e"></a>

## Next pages — items.object.system_metadata.owner_view / ca96e2798127 / 8

- [items.object.system_metadata](data-sources--site_registrations_by_state--reference--group-002.md#canonical-877448804fbc002b6806788021a0e670d7179f38b80a98f00469b87c1f686f82)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-9f1c2c7bdb366f4d4e91f49c44a8532f1c3abe65a9944ef6d3941b5c1acd7aba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8df7efe3c94d5c00616319eb83067b03e1a2f0da484dc4f024f66bcc18e4a6b6"></a>

## items.owner_view — items.owner_view / 5468ffe86d53 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- items.owner_view

<a id="canonical-dcadc76bdff1c3b3510abe16b29d68cd606e60675015d4e6429ac5d7170afd6e"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

<a id="canonical-0791952eec2149b0113a46032f338296f1b416f641ad12bc90447f9528fd5339"></a>

## Direct properties — items.owner_view / 5468ffe86d53 / 3

<a id="canonical-640b9955805dd2e29f2759bd4d4a9472f448ddf8da38ed20c837ff8d62b43438"></a>

<a id="canonical-0ffca0534b741a344fd1d2e881d074d8cd23e2235f1715589d5b1354e7495a30"></a>

## kind property — items.owner_view / 5468ffe86d53 / 4

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="canonical-015b1f4361b04d55fd582da935c3095b111d75dd5e60b35af6dad073a4d2cdf8"></a>

<a id="canonical-0a89220293c34d2da2000743baf653a615f760099da0febef3e01a46bd1ae9b1"></a>

## name property — items.owner_view / 5468ffe86d53 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-6dbd08b5c1512c8d7b8ebabfd6f2779ab3cd054d7a7d0548230c17f40c5dc9d6"></a>

<a id="canonical-1860766f776db32bf0968e190887dec24962e15abbe6ca98a65622429b1f2e20"></a>

## namespace property — items.owner_view / 5468ffe86d53 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-bf6b7fe8a20fb7b50e198980d8475e7de5a7729cea6b2775932633c9c901ebe0"></a>

<a id="canonical-84a6cdb416275dd8fa153f6c519ce94e456372adbab453f93e958fbe2d9bf8c7"></a>

## uid property — items.owner_view / 5468ffe86d53 / 7

Type: `"string"`. Computed.

UID. UID of the view object.

<a id="canonical-b64a226fa3f975fe2bb4084cc9cce7e3f8f82ca792fbbb62699855358040a28e"></a>

## Next pages — items.owner_view / 5468ffe86d53 / 8

- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8bc88403024cae7ca7326d87387a53f6c91e90b0e4fbdbdd358ecb2a41db2b14"></a>

## items.system_metadata — items.system_metadata / da229eac8306 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- items.system_metadata

<a id="canonical-fbf9d31e1277cfb4777d53ddcf1be885d2271d7c1d6873cfe06b9ace35387688"></a>

Type: `"single"`. Computed.

SystemObjectGetMetaType is metadata generated or populated by the system for all persisted objects
and cannot be updated directly by users.

<a id="canonical-8d8737877444d565a829743410d7c72ef119495294748ed7a993bae6cdd90f60"></a>

## Direct properties — items.system_metadata / da229eac8306 / 3

<a id="canonical-d156ad066a0335215b33bf80d1039c2fd3ba416ddd62e8a04b37e7326ca96e19"></a>

<a id="canonical-1e6f2bd744778927915fd1ebd8063944e7044b457d98e8dea7eb8cc49c8b933e"></a>

## creation_timestamp property — items.system_metadata / da229eac8306 / 4

Type: `"string"`. Computed.

CreationTimestamp is a timestamp representing the server time when this object was created. It is
not guaranteed to be set in happens-before order across separate operations. Clients may not set
this value.

<a id="canonical-dbe7e65fcbc5b82da466da72e29132f434e8308f803ce985eb1c1f12751a075c"></a>

<a id="canonical-c2735041e7951e9bd67a6665ca940cf6c4be114f0cc44ea781327f52ffedc543"></a>

## creator_class property — items.system_metadata / da229eac8306 / 5

Type: `"string"`. Computed.

Value identifying the class of the user or service which created this configuration object.

<a id="canonical-31eecb18d371edf103fa8bd561feafe6eb3968cd278bc36f814173111f4bf64a"></a>

<a id="canonical-5053e61b252058e80bcdf3b50f1b29e72dfebb57e925cc5f67be8ce4f247c55f"></a>

## creator_id property — items.system_metadata / da229eac8306 / 6

Type: `"string"`. Computed.

Value identifying the exact user or service that created this configuration object.

<a id="canonical-f9689973d7ab11ea22e93b94c2fd69fd4cdccaca9091e8328a4b20bc49b413f3"></a>

<a id="canonical-e98652c6e26eaec1eef87160a54695099f8d3969c75bfdafcce7e7767b242451"></a>

## deletion_timestamp property — items.system_metadata / da229eac8306 / 7

Type: `"string"`. Computed.

DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is
set by the server when a graceful deletion is requested by the user, and is not directly settable by
a client. The resource is expected to be deleted (no longer visible from resource lists, and not..

<a id="canonical-eb733c20f3e48fb3d741c6392f74dd8c447f841909fdeb37a73e1437bd9278ec"></a>

<a id="canonical-405290743db4e857134d6abe8d60f36f7a82336ba7469c27a53eeb189bd681e7"></a>

## finalizers property — items.system_metadata / da229eac8306 / 8

Type: `["list", "string"]`. Computed.

Must be empty before the object is deleted from the registry. Each entry is an identifier for the
responsible component that will remove the entry from the list. If the deletionTimestamp of the
object is non-nil, entries in this list can only be removed.

- [initializers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-7c1d718575f0bb7edd02972bb91c0ff8af9caa7eaa323a88b10c31b63b8bf30e): complete subsection reference.

- [labels](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1c1c8839a4e412bc475a708b5cb8a5420d7c64a94ffbcc78cd42fd10a5d5eee6): complete subsection reference.

<a id="canonical-068be60826a873b7fbd5c56770c232e2f6f4dcc31c4459fbb75dc5026144d680"></a>

<a id="canonical-00fbb25bcc00a6a2e30678a79bda6455b47e27b79aaa7d65aa157dd64dc8c49a"></a>

## modification_timestamp property — items.system_metadata / da229eac8306 / 9

Type: `"string"`. Computed.

ModificationTimestamp is a timestamp representing the server time when this object was last
modified.

<a id="canonical-9a1ae02dc798c9a59e0ef36c1e21b24bdf098dcd2856c7b82982b5fde28df68c"></a>

<a id="canonical-8e76ba036a4553645afb56051e0bedce1dea169295f47aa70f2261b4cc312f42"></a>

## object_index property — items.system_metadata / da229eac8306 / 10

Type: `"number"`. Computed.

Unique index for the object. Some objects need a unique integer index to be allocated for each
object type. This field will be populated for all objects that need it and will be zero otherwise.

- [owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-7cbbd1175e2727f4299f7b5d62407786b1d56149eae6f69109f918b6a8104afb): complete subsection reference.

<a id="canonical-56370735295d422a9ec381396cd277a276d4f02bce4771a18462850c9a72362f"></a>

<a id="canonical-54f4a1f33bdba6fe5ac1d2bd066846e50904753a27075f964cc27fd17b328c2d"></a>

## tenant property — items.system_metadata / da229eac8306 / 11

Type: `"string"`. Computed.

Tenant to which this configuration object belongs to. The value for this is found from presented
credentials.

<a id="canonical-3e661322f5688ff710c516107d8b4a38244c7ec7de0ed8c5e89de8072f2ac9ef"></a>

<a id="canonical-235117d40a0630238ebeb8052723719f9f220e4f31447678facd94a63b57cbae"></a>

## uid property — items.system_metadata / da229eac8306 / 12

Type: `"string"`. Computed.

Uid is the unique in time and space value for this object. It is generated by the server on
successful creation of an object and is not allowed to change on Replace API. The value of is taken
from uid field of ObjectMetaType, if provided.

<a id="canonical-e553383005057d53a1bd1a1f7770e310672c06fbbe9cf0e01418ca711a45db8d"></a>

## Next pages — items.system_metadata / da229eac8306 / 13

- [items.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-7c1d718575f0bb7edd02972bb91c0ff8af9caa7eaa323a88b10c31b63b8bf30e)
- [items.system_metadata.labels](data-sources--site_registrations_by_state--reference--group-003.md#canonical-1c1c8839a4e412bc475a708b5cb8a5420d7c64a94ffbcc78cd42fd10a5d5eee6)
- [items.system_metadata.owner_view](data-sources--site_registrations_by_state--reference--group-003.md#canonical-7cbbd1175e2727f4299f7b5d62407786b1d56149eae6f69109f918b6a8104afb)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-7c1d718575f0bb7edd02972bb91c0ff8af9caa7eaa323a88b10c31b63b8bf30e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd8c5af2f493bcf1430cf4d297707895ea9706761181d9372d392f0475f27334"></a>

## items.system_metadata.initializers — items.system_metadata.initializers / d2e3f86e0e48 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba)
- items.system_metadata.initializers

<a id="canonical-e0fff57ec3f45ae35f2ff8895289bffc542ae4e9bd38192d04fde852833387ee"></a>

Type: `"single"`. Computed.

Initializers tracks the progress of initialization of a configuration object.

<a id="canonical-4fc758d30365827ce9f3ceadd95dd7148964a89a26eb53097f047bb77c917acc"></a>

## Direct properties — items.system_metadata.initializers / d2e3f86e0e48 / 3

- [pending](data-sources--site_registrations_by_state--reference--group-003.md#canonical-a74f18f1271c495d147c8e908746c5201fc21d6fc389bd940bb97c5d27178d23): complete subsection reference.

- [result](data-sources--site_registrations_by_state--reference--group-003.md#canonical-104b12cfa6047b3b60b478963739f6cfce17867e6e842da496cd41ecbb8c8836): complete subsection reference.

<a id="canonical-c631460eaec93e0abee55c76a611a836256ab061fd6f35392cd2d73554fd70a7"></a>

## Next pages — items.system_metadata.initializers / d2e3f86e0e48 / 4

- [items.system_metadata.initializers.pending](data-sources--site_registrations_by_state--reference--group-003.md#canonical-a74f18f1271c495d147c8e908746c5201fc21d6fc389bd940bb97c5d27178d23)
- [items.system_metadata.initializers.result](data-sources--site_registrations_by_state--reference--group-003.md#canonical-104b12cfa6047b3b60b478963739f6cfce17867e6e842da496cd41ecbb8c8836)
- [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-a74f18f1271c495d147c8e908746c5201fc21d6fc389bd940bb97c5d27178d23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f12b9a034a7efc6f05b41808c1cf3efbc0c2b83e366613dd81ef0f48c3cb799c"></a>

## items.system_metadata.initializers.pending — items.system_metadata.initializers.pending / 2404c4dc2483 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba)
- [items.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-7c1d718575f0bb7edd02972bb91c0ff8af9caa7eaa323a88b10c31b63b8bf30e)
- items.system_metadata.initializers.pending

<a id="canonical-2626063209967fe3bda32b6525355cac5cb7e8436871d792cc53a7edae3b19e7"></a>

Type: `"list"`. Computed.

Pending is a list of initializers that must execute in order before this object is initialized. When
the last pending initializer is removed, and no failing result is set, the initializers struct will
be set to nil and the object is considered as initialized and visible to all clients.

<a id="canonical-9435485fffbd4099ad02a227da66394e7555219176495dd95351046b8c14933e"></a>

## Direct properties — items.system_metadata.initializers.pending / 2404c4dc2483 / 3

<a id="canonical-55b8fab150b784817cf4d166129a44147b9e5990879c160d1e4410ec648ce8c4"></a>

<a id="canonical-d045dfddb9cf5890fa93a35864b7c7f51d3689f9ed82b367d1ad38fb7eff6d5c"></a>

## name property — items.system_metadata.initializers.pending / 2404c4dc2483 / 4

Type: `"string"`. Computed.

Name of the service that is responsible for initializing this object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-5a068644efe2e4b79e3b0eedf45adfda7453b30e2cbfd52b8585c96601ca49f1"></a>

## Next pages — items.system_metadata.initializers.pending / 2404c4dc2483 / 5

- [items.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-7c1d718575f0bb7edd02972bb91c0ff8af9caa7eaa323a88b10c31b63b8bf30e)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-104b12cfa6047b3b60b478963739f6cfce17867e6e842da496cd41ecbb8c8836"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d31b64875122733c883624104b2102cf15f05e8befc814d79df1aa12d9f2847"></a>

## items.system_metadata.initializers.result — items.system_metadata.initializers.result / 1157ab669952 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba)
- [items.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-7c1d718575f0bb7edd02972bb91c0ff8af9caa7eaa323a88b10c31b63b8bf30e)
- items.system_metadata.initializers.result

<a id="canonical-e764da52e50848efbb1ab52820befcd88ebccb09416e8928f76853957692735e"></a>

Type: `"single"`. Computed.

Status is a return value for calls that don't return other objects.

<a id="canonical-843073392a6e1c8b90e8350c411d727e5c314a894518bd0c4fdcd7fcb7c178ed"></a>

## Direct properties — items.system_metadata.initializers.result / 1157ab669952 / 3

<a id="canonical-4f1dec2668141d20df129d4d9363dc6ee685e87cc3363557bb322ba6b9820364"></a>

<a id="canonical-4a6d723bb42c7b9ca2279258cbe84659facacdc0d6da4d724815e8033eb72ddf"></a>

## code property — items.system_metadata.initializers.result / 1157ab669952 / 4

Type: `"number"`. Computed.

Suggested HTTP return code for this status, 0 if not set.

<a id="canonical-5005caf20fa9de05ce64fc2daff8d73c763b099b8f0a75d98af0b44fd55eb0d7"></a>

<a id="canonical-0a357e7dcaac13b32663be2de0036bcc9fbaa3ea9323ea3e68a1ea9cc91af71f"></a>

## reason property — items.system_metadata.initializers.result / 1157ab669952 / 5

Type: `"string"`. Computed.

Human-readable description of why this operation is in the 'Failure' status. If this value is empty
there is no information available.

<a id="canonical-e44f62feaddb4435569642afbf717cf04e6078e03602809e51e507532f566f89"></a>

<a id="canonical-ec174c414053e5259ede7955a333e9e3ecbf03fd189b4f15e12f6959001b4f10"></a>

## status property — items.system_metadata.initializers.result / 1157ab669952 / 6

Type: `"string"`. Computed.

Status of the operation. One of: 'Success' or 'Failure'.

<a id="canonical-ff35e456d081f0b90486abdbb2381b0b95ab40952254ddcb2053c31ad0053785"></a>

## Next pages — items.system_metadata.initializers.result / 1157ab669952 / 7

- [items.system_metadata.initializers](data-sources--site_registrations_by_state--reference--group-003.md#canonical-7c1d718575f0bb7edd02972bb91c0ff8af9caa7eaa323a88b10c31b63b8bf30e)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-1c1c8839a4e412bc475a708b5cb8a5420d7c64a94ffbcc78cd42fd10a5d5eee6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ad8bed9a9d48e5f7d41c1072de9bba9fb2cdf86b8468512c23c066c0d6c84aa"></a>

## items.system_metadata.labels — items.system_metadata.labels / a0c052d3b052 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba)
- items.system_metadata.labels

<a id="canonical-981a2260cf39b33bf4529ea1f3d45d1b012c27f195877640d12befcf586f0c2e"></a>

Type: `"single"`. Computed.

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the operator or software. Values here can be interpreted by software(backend or
frontend) to enable certain behavior e.g. Things marked as soft-deleted(restorable).

<a id="canonical-59205ea2870513a2c7a6c1b25d901d2189976df6464f2d3db3656989a99c97ba"></a>

## Direct properties — items.system_metadata.labels / a0c052d3b052 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d3bbb4470026f43701b199cc512940783ed2586a492828c2dfc0119c35268e7a"></a>

## Next pages — items.system_metadata.labels / a0c052d3b052 / 4

- [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)

<a id="canonical-7cbbd1175e2727f4299f7b5d62407786b1d56149eae6f69109f918b6a8104afb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9fa0a5cf4eeae1bbec94518f8de5ee99c51677f15bed3f7529f9d4345323762"></a>

## items.system_metadata.owner_view — items.system_metadata.owner_view / fc9476e858a0 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
- [Property reference](data-sources--site_registrations_by_state--reference--group-001.md#canonical-38167b71523ef96832758bdc9a667a238f95e7a53fb77a57e4c1b2c9896a9b1c)
- [items](data-sources--site_registrations_by_state--reference--group-001.md#canonical-5f68d9c7b6bc367983315aa81b9bf692573d859a7cbaaf2f55ceb4eb8801c66c)
- [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba)
- items.system_metadata.owner_view

<a id="canonical-2e76e92ce2c67503c6ee58b6c4cbeabe2ad804ad4b3f8cca7f37958f021d6d14"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

<a id="canonical-0569d4d3c80099e4532c1b498843abee7216827c94406dd4b45ebc694ebcfd3c"></a>

## Direct properties — items.system_metadata.owner_view / fc9476e858a0 / 3

<a id="canonical-dbb710c2e186beed0d35fe1eb8cf21ad44ee884736e4a6a4171139922063c757"></a>

<a id="canonical-d8cef537d3f283b29575edd0b1136655980a4ddf20bcfd19974ef6ae6c1a39b3"></a>

## kind property — items.system_metadata.owner_view / fc9476e858a0 / 4

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="canonical-5ea5c115c0dcec04c31213bbc82ba79d6ddc006e182127897e77770f2c09626c"></a>

<a id="canonical-5ff47ae842783bb5dc7d4956313edf5fc35690a43edc0b63adb7e66f68265e4c"></a>

## name property — items.system_metadata.owner_view / fc9476e858a0 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-f3291aa7e2891010d68e1b5fa64d5c96de9297c623f90f14d8552ae3760a97d5"></a>

<a id="canonical-561b1fa1b2c7fcae38a38c8e1476d9047b276a46ad724d65748ec7926dfcf04f"></a>

## namespace property — items.system_metadata.owner_view / fc9476e858a0 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="canonical-240766e3f3c4d11552326242b526c1d11956e483ecd23eda4cc01736c9bc7dd9"></a>

<a id="canonical-ae468017fbc2bfe52c14856e21bd02cc4c4a5de5f46e111135ff595eabc9a034"></a>

## uid property — items.system_metadata.owner_view / fc9476e858a0 / 7

Type: `"string"`. Computed.

UID. UID of the view object.

<a id="canonical-e4cbcdc6987bdfbd772e3b86305c26aaab8b6954c375ab88ae5910bb0134b2e0"></a>

## Next pages — items.system_metadata.owner_view / fc9476e858a0 / 8

- [items.system_metadata](data-sources--site_registrations_by_state--reference--group-003.md#canonical-3ae60735a1fec18d77c2467983d01bce022e79f2c5e865874b8c735a149673ba)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md#canonical-45d2eb09f36f60d5d82eb56138e95138ce81ee60f4449a9d1839dc3d217ebdf3)
