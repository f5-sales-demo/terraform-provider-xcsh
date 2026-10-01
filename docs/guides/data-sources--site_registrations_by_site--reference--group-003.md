---
page_title: "xcsh_site_registrations_by_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_site reference."
---

# xcsh_site_registrations_by_site reference

<a id="canonical-517b7097b2287aff701cb3f3307910dc21e9d27a313aa5cdd62ead840093bfc1"></a>

## Next pages — items.object.system_metadata.initializers / 07189ef39884 / 4

- [items.object.system_metadata.initializers.pending](data-sources--site_registrations_by_site--reference--group-003.md#canonical-376c0295894dc343ecf52fd6bc8159bc02150795ae917299e73e7b4f53e9f1cc)
- [items.object.system_metadata.initializers.result](data-sources--site_registrations_by_site--reference--group-003.md#canonical-db9f7c67e6a178cf219b5ce6375e011de767065d9f141b22589262740fa9a83c)
- [items.object.system_metadata](data-sources--site_registrations_by_site--reference--group-002.md#canonical-2eb083250b6320ac77d7d040c45a97995c3731f2f3867e06533561c20ab517de)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-376c0295894dc343ecf52fd6bc8159bc02150795ae917299e73e7b4f53e9f1cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-567203b989a42fd09c0e727e3a1a99c441330b4445c56d38000cc38718ae2438"></a>

## items.object.system_metadata.initializers.pending — items.object.system_metadata.initializers.pending / 83d9d8f2a89a / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.object](data-sources--site_registrations_by_site--reference--group-002.md#canonical-0dbf9c88ad096a90aef7269fa26105ad2067f1a68b64d6a07e96a9f90bdebe33)
- [items.object.system_metadata](data-sources--site_registrations_by_site--reference--group-002.md#canonical-2eb083250b6320ac77d7d040c45a97995c3731f2f3867e06533561c20ab517de)
- [items.object.system_metadata.initializers](data-sources--site_registrations_by_site--reference--group-002.md#canonical-029ffe0a7cb2d080a16fcad872db70ffb5c39508afc7192f540bbb8f24b228f5)
- items.object.system_metadata.initializers.pending

<a id="canonical-9c386e2af99b77bfed54a3dfac92e931d204e883dca366f63ee5379d8993769b"></a>

Type: `"list"`. Computed.

Pending is a list of initializers that must execute in order before this object is initialized. When
the last pending initializer is removed, and no failing result is set, the initializers struct will
be set to nil and the object is considered as initialized and visible to all clients.

<a id="canonical-cb89436e9c7f34b54e766df1fc295e3fb9bcff2a96586e03a6488fe9f9f85425"></a>

## Direct properties — items.object.system_metadata.initializers.pending / 83d9d8f2a89a / 3

<a id="canonical-2328cd87fc99e538b96d253a52d60e6b72fafeba7b8320e9c22404e1753589b0"></a>

<a id="canonical-67ff2c6df01eab21047bc411a822e3104586870d5bb484d7f114f74125c88642"></a>

## name property — items.object.system_metadata.initializers.pending / 83d9d8f2a89a / 4

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

<a id="canonical-f33209cbf669a85b4032000816ba256386bec35fc20692906d4b2f2fa304c7cd"></a>

## Next pages — items.object.system_metadata.initializers.pending / 83d9d8f2a89a / 5

- [items.object.system_metadata.initializers](data-sources--site_registrations_by_site--reference--group-002.md#canonical-029ffe0a7cb2d080a16fcad872db70ffb5c39508afc7192f540bbb8f24b228f5)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-db9f7c67e6a178cf219b5ce6375e011de767065d9f141b22589262740fa9a83c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2374e10c412399b0faf921fbe266a277bf5a58e8d00974a636d592f871c4c084"></a>

## items.object.system_metadata.initializers.result — items.object.system_metadata.initializers.result / d1b14bfb7445 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.object](data-sources--site_registrations_by_site--reference--group-002.md#canonical-0dbf9c88ad096a90aef7269fa26105ad2067f1a68b64d6a07e96a9f90bdebe33)
- [items.object.system_metadata](data-sources--site_registrations_by_site--reference--group-002.md#canonical-2eb083250b6320ac77d7d040c45a97995c3731f2f3867e06533561c20ab517de)
- [items.object.system_metadata.initializers](data-sources--site_registrations_by_site--reference--group-002.md#canonical-029ffe0a7cb2d080a16fcad872db70ffb5c39508afc7192f540bbb8f24b228f5)
- items.object.system_metadata.initializers.result

<a id="canonical-d4ba1b855a8a0f5d75b7a3ef33e4ebe36cd79d1d5bc83fc95efe5cc3e89cf82c"></a>

Type: `"single"`. Computed.

Status is a return value for calls that don't return other objects.

<a id="canonical-ded5bda954189df328e1d752a28f33cc11467ec7dc87d0c92eda35fe95dd7126"></a>

## Direct properties — items.object.system_metadata.initializers.result / d1b14bfb7445 / 3

<a id="canonical-23a04908d3d70973ef0d6941da2056ba8e630a9421e9f1c48fd43c6485d9d2d9"></a>

<a id="canonical-5a46484c4a0492767f0130ada93c34cf66d557b896b09a13b3574f353544e3db"></a>

## code property — items.object.system_metadata.initializers.result / d1b14bfb7445 / 4

Type: `"number"`. Computed.

Suggested HTTP return code for this status, 0 if not set.

<a id="canonical-d4d998ee4316161045792b68348d7809ceba5500513d47ef876e6937282e1003"></a>

<a id="canonical-322579c01f4e8bcc8ea1e7c0c790c2bc1678e06d3dca531b704550a19bc8b2f7"></a>

## reason property — items.object.system_metadata.initializers.result / d1b14bfb7445 / 5

Type: `"string"`. Computed.

Human-readable description of why this operation is in the 'Failure' status. If this value is empty
there is no information available.

<a id="canonical-0e05f87b1352f7b860a513247c37adc6b5bb4e34ac9414517b836fc50c641448"></a>

<a id="canonical-904af9726c5fd88dd92aeff5fbda4850c1d2fba343ec9358745f6812d2bf77c7"></a>

## status property — items.object.system_metadata.initializers.result / d1b14bfb7445 / 6

Type: `"string"`. Computed.

Status of the operation. One of: 'Success' or 'Failure'.

<a id="canonical-fd9a15c00bda2731c398a85f1ae507457cd457c421c2dc982d53a26f84996fbc"></a>

## Next pages — items.object.system_metadata.initializers.result / d1b14bfb7445 / 7

- [items.object.system_metadata.initializers](data-sources--site_registrations_by_site--reference--group-002.md#canonical-029ffe0a7cb2d080a16fcad872db70ffb5c39508afc7192f540bbb8f24b228f5)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-e1db626d8ab0ba7ba196c173a1429a8cc4a21ef715a9a37efe2ef0b2af657ba1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d634ba77b4f9c78388dffb9a489dabe72b0171da92f32410a5210267d3c4ffc"></a>

## items.object.system_metadata.labels — items.object.system_metadata.labels / b2b4210e8df8 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.object](data-sources--site_registrations_by_site--reference--group-002.md#canonical-0dbf9c88ad096a90aef7269fa26105ad2067f1a68b64d6a07e96a9f90bdebe33)
- [items.object.system_metadata](data-sources--site_registrations_by_site--reference--group-002.md#canonical-2eb083250b6320ac77d7d040c45a97995c3731f2f3867e06533561c20ab517de)
- items.object.system_metadata.labels

<a id="canonical-f5a072907bed4e4f65810902c07f16e6dc2c16ed0f50f285e1b568612915ead7"></a>

Type: `"single"`. Computed.

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the operator or software. Values here can be interpreted by software(backend or
frontend) to enable certain behavior e.g. Things marked as soft-deleted(restorable).

<a id="canonical-a457258d002cfa623ca297f9ba4a4c00c30511dd3f62ece7d4d6c317e3824d23"></a>

## Direct properties — items.object.system_metadata.labels / b2b4210e8df8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-446d61665a8edc67500f248f284701fd04883131ed27ec65893c332bff44b8b4"></a>

## Next pages — items.object.system_metadata.labels / b2b4210e8df8 / 4

- [items.object.system_metadata](data-sources--site_registrations_by_site--reference--group-002.md#canonical-2eb083250b6320ac77d7d040c45a97995c3731f2f3867e06533561c20ab517de)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-d4c84e4e1b4bf139e1d17dd308f4b2114af8a4b1a59080c7fb75919fd1d8a785"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca5ac7c951f6a07d399c88eb18615e73fcefdcaf1eefb75033b9c5ca025e698e"></a>

## items.object.system_metadata.namespace — items.object.system_metadata.namespace / 3a9b8d232a37 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.object](data-sources--site_registrations_by_site--reference--group-002.md#canonical-0dbf9c88ad096a90aef7269fa26105ad2067f1a68b64d6a07e96a9f90bdebe33)
- [items.object.system_metadata](data-sources--site_registrations_by_site--reference--group-002.md#canonical-2eb083250b6320ac77d7d040c45a97995c3731f2f3867e06533561c20ab517de)
- items.object.system_metadata.namespace

<a id="canonical-190e805c01329200b38d85e9ae130becd8cc546234d7862a92fc1f42840847fb"></a>

Type: `"list"`. Computed.

The namespace this object belongs to. This is populated by the service based on the
metadata.namespace field when an object is created.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

<a id="canonical-59358baddc14ab2a4a38feca1dd0ce6a6548d9223ec3846a5484b65f4ab5fd2d"></a>

## Direct properties — items.object.system_metadata.namespace / 3a9b8d232a37 / 3

<a id="canonical-3250be105fddafc253a053c8b35cac6592ba1d837372beb5c28b85b15c416b55"></a>

<a id="canonical-85c04a358dc743aaaffe9dded88cf3037d3b8201fa3b029099aa28d415957f3f"></a>

## kind property — items.object.system_metadata.namespace / 3a9b8d232a37 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="canonical-3d7c63b565db4605fde713f16108d84b4796ec92dd0b342236031825ce2e39dd"></a>

<a id="canonical-6252c46eb78e4c9a592119db37bac7cfa3d8222722c404c33356660e0164e4eb"></a>

## name property — items.object.system_metadata.namespace / 3a9b8d232a37 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="canonical-9ca6ad9593027004844a9ff31caaf9f10987cbe9ab4d9fd0216f0759eed40b38"></a>

<a id="canonical-157befd33c629e87b44dab176d11969d8b84132d778e1ad7f745b2bf27d26342"></a>

## namespace property — items.object.system_metadata.namespace / 3a9b8d232a37 / 6

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

<a id="canonical-0c5b6c6c541e6e733a8f2be12b5440c061abc2417e2acb4d15a5c3ea29db83c6"></a>

<a id="canonical-330f135a625b0587ee0f0b63f2431f03f2dbc7a42cff53ecff90084eede50242"></a>

## tenant property — items.object.system_metadata.namespace / 3a9b8d232a37 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="canonical-3d85cd9430211879cee80695a8dd06a7e5b3b87e9fb946c06f7b7b0d69e4fddd"></a>

<a id="canonical-fa82f308566c0e24d6ece1488046c4bf502b3e52eff63cc2d8c9fb0f9fa108c2"></a>

## uid property — items.object.system_metadata.namespace / 3a9b8d232a37 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

<a id="canonical-a8fb90f867042b45b9883f8f57b5f7dac015cb2d6b82e749b56103563481f341"></a>

## Next pages — items.object.system_metadata.namespace / 3a9b8d232a37 / 9

- [items.object.system_metadata](data-sources--site_registrations_by_site--reference--group-002.md#canonical-2eb083250b6320ac77d7d040c45a97995c3731f2f3867e06533561c20ab517de)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-e6899d7867e5040f5b4cb55f242bbfc2354830de7364170632e825f713060c4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c85d50124608eb6cfbc71336446a25893d38edc3f4411a2a6191cacb07650d2e"></a>

## items.object.system_metadata.owner_view — items.object.system_metadata.owner_view / 75bf7d6a1efc / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.object](data-sources--site_registrations_by_site--reference--group-002.md#canonical-0dbf9c88ad096a90aef7269fa26105ad2067f1a68b64d6a07e96a9f90bdebe33)
- [items.object.system_metadata](data-sources--site_registrations_by_site--reference--group-002.md#canonical-2eb083250b6320ac77d7d040c45a97995c3731f2f3867e06533561c20ab517de)
- items.object.system_metadata.owner_view

<a id="canonical-2efc4627e2fadf2203521376fdbd8a22b52332c931cdc1cff36d5de52555f7d3"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

<a id="canonical-141801d5c3c277e82f417f535d7d1a2818ca35840465e6f04fd1068b1f1bcd8f"></a>

## Direct properties — items.object.system_metadata.owner_view / 75bf7d6a1efc / 3

<a id="canonical-c27526eb4d49c11f25d9044096a18313425be97ed0f210da0d0f8822bacbd1ce"></a>

<a id="canonical-67df865bdfee7f218fcc1686dd334cb95740b95d68413cf348664f259e1f041a"></a>

## kind property — items.object.system_metadata.owner_view / 75bf7d6a1efc / 4

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="canonical-28cd66de2a1c47b99d18e936c84f4fbbfaf91ceee6829415cfc05cba6cf45e2d"></a>

<a id="canonical-72b67581831b7cd56da42f38703c1b69ebe87cfe03668e0d5439f58137e94294"></a>

## name property — items.object.system_metadata.owner_view / 75bf7d6a1efc / 5

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

<a id="canonical-2373a45e2d86cbce84ea1ed0fc8171e15e68be3da434e70d7f3aeecbac5ce60e"></a>

<a id="canonical-1e12216fe65da396483da87ffb7347dd01b1e63ad6ae29c521f40a1e13966bc7"></a>

## namespace property — items.object.system_metadata.owner_view / 75bf7d6a1efc / 6

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

<a id="canonical-a73a1cf54b95526ef1b514e46ef50e0a5815a04d14708538a30caa6394cc11a3"></a>

<a id="canonical-c5ce618ecc4f295d80f78808935afbb5c83e63dae90bafe7a65b613514523256"></a>

## uid property — items.object.system_metadata.owner_view / 75bf7d6a1efc / 7

Type: `"string"`. Computed.

UID. UID of the view object.

<a id="canonical-c05b59d1e6f61e61c6f429d34491f98f5c43a6c29306fff27656dbd89f3b8662"></a>

## Next pages — items.object.system_metadata.owner_view / 75bf7d6a1efc / 8

- [items.object.system_metadata](data-sources--site_registrations_by_site--reference--group-002.md#canonical-2eb083250b6320ac77d7d040c45a97995c3731f2f3867e06533561c20ab517de)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-889ff3191d71bb3551f5930fc0915299b3c277568ecfc630ac9a31e4871badaa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-252df7bfef4a4aafecf822d6faf957ede63bd862ceccdde5db85c9f385374c11"></a>

## items.owner_view — items.owner_view / c52ba08c75f7 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- items.owner_view

<a id="canonical-553fca4c87652bc4f6b7781455a2ac206dafaebe71de98312e37067b5bb9644f"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

<a id="canonical-8428851fde8c43283b6eb67afdd517758adeacce604d17c58982700d90eb2519"></a>

## Direct properties — items.owner_view / c52ba08c75f7 / 3

<a id="canonical-b4cf4cf590a1aa9b18bbbdcbf19d251cd8093e1a026d0cf50ca6910f6918f372"></a>

<a id="canonical-ee745e29fbfd3f26af267d9112c518993c020f7474e0887ce3c3ace854040d79"></a>

## kind property — items.owner_view / c52ba08c75f7 / 4

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="canonical-40473320862385cc2de4019c5c6c86c2ae2ab9dd59c8490347fd1b9f003b648e"></a>

<a id="canonical-ac82335b9db56a8bcc6099c447ae4416b0454153a85fbd1abbc1aecac01ebe7c"></a>

## name property — items.owner_view / c52ba08c75f7 / 5

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

<a id="canonical-83f6b101ad593dd5f4167ff9e6f901bfce021cb34d68ed3eac5a04c332b1a52f"></a>

<a id="canonical-252fcae244c295758fb1eed86631884bd75ee4c260ed3ee658270a2cff5fae04"></a>

## namespace property — items.owner_view / c52ba08c75f7 / 6

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

<a id="canonical-199e0d665debb83acee873114b0f431aaf3cd457cfc64ef45530e959e16442d2"></a>

<a id="canonical-3654b414913899a0fe066515b9eb4e25a5d510945237ce92577d6f4a26c73130"></a>

## uid property — items.owner_view / c52ba08c75f7 / 7

Type: `"string"`. Computed.

UID. UID of the view object.

<a id="canonical-2c13389877ce192a8bd3f2597a102df024f888ab79306dfd1f25cd7c4980b8fc"></a>

## Next pages — items.owner_view / c52ba08c75f7 / 8

- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-ad1a56f7a17b1bf18c3d85dc22c426ffab344a75a683a433efbf413a7d2b84fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66fcd36e938700f40cf79195eac0ec82333603a82a77b1abbcf6ddf12a24f357"></a>

## items.system_metadata — items.system_metadata / e9f1122c1d84 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- items.system_metadata

<a id="canonical-5a2a1e20a5f95059850720fd7f37b0ecd5b94b2fc879e2e16526a4670e994a4c"></a>

Type: `"single"`. Computed.

SystemObjectGetMetaType is metadata generated or populated by the system for all persisted objects
and cannot be updated directly by users.

<a id="canonical-ee31f64e258527f42997529d3f0775094bcff85ac55e47a5ae75ceae6d5d5a55"></a>

## Direct properties — items.system_metadata / e9f1122c1d84 / 3

<a id="canonical-a2fe4ad8abbc4f1001143f905415c27dd4da1679a103bdfe3a1df172b0411cbe"></a>

<a id="canonical-8c3e9909bef478e9efdc28912d128d722c627052d020142dc269d09d03fa3533"></a>

## creation_timestamp property — items.system_metadata / e9f1122c1d84 / 4

Type: `"string"`. Computed.

CreationTimestamp is a timestamp representing the server time when this object was created. It is
not guaranteed to be set in happens-before order across separate operations. Clients may not set
this value.

<a id="canonical-e73e265026f52b06c3ade51de1ba3731b0d4227435ca3b355924db416b76a495"></a>

<a id="canonical-79de9a5e574a9408eda15bb3aad80574f312548b05c49f876a9d950ba89bd452"></a>

## creator_class property — items.system_metadata / e9f1122c1d84 / 5

Type: `"string"`. Computed.

Value identifying the class of the user or service which created this configuration object.

<a id="canonical-640c9ce4fc233658ce663ba972021cf2fdcebe69bccf12b711dd10fb80b4aa05"></a>

<a id="canonical-88b698d69a664b54b3a603248a7079883affd0c2e9a399f5ff34793eb891488b"></a>

## creator_id property — items.system_metadata / e9f1122c1d84 / 6

Type: `"string"`. Computed.

Value identifying the exact user or service that created this configuration object.

<a id="canonical-33647848dbd3198e09ec5d283f516599d03ceb94681e8440648b076de4cba36c"></a>

<a id="canonical-854ba4a28c00fd7db6e53e98bb368aff69a89c0d7c068d982c5db1fd3c703e9e"></a>

## deletion_timestamp property — items.system_metadata / e9f1122c1d84 / 7

Type: `"string"`. Computed.

DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is
set by the server when a graceful deletion is requested by the user, and is not directly settable by
a client. The resource is expected to be deleted (no longer visible from resource lists, and not..

<a id="canonical-0e14b20c077cf1b09c2b1fb498d5aae4ed0cfdc91262ccaa96745a7a60e7b2e0"></a>

<a id="canonical-2765c51f48ba9da7157228146a9cb45269a38a05a42c68d07fd4116b1988b288"></a>

## finalizers property — items.system_metadata / e9f1122c1d84 / 8

Type: `["list", "string"]`. Computed.

Must be empty before the object is deleted from the registry. Each entry is an identifier for the
responsible component that will remove the entry from the list. If the deletionTimestamp of the
object is non-nil, entries in this list can only be removed.

- [initializers](data-sources--site_registrations_by_site--reference--group-003.md#canonical-d38e9f13f5a9efa85a0ea7e1a59fc2a4d1d976f4663ee4f28f090d4035bd744e): complete subsection reference.

- [labels](data-sources--site_registrations_by_site--reference--group-003.md#canonical-658639cc2e573c337edccd92785970404e32e104638d452e8245ec3e7b2193f3): complete subsection reference.

<a id="canonical-d396d6b7b5fbd9e5ea3f7ef0fcbf10808f1dceb87b919229e035039a55267625"></a>

<a id="canonical-456aa7d8fdc865e777c849e0d55f471e1d30894876c20a885e26deed513002d3"></a>

## modification_timestamp property — items.system_metadata / e9f1122c1d84 / 9

Type: `"string"`. Computed.

ModificationTimestamp is a timestamp representing the server time when this object was last
modified.

<a id="canonical-35cc02cbe8cc074d608561e44b681deaec94117d063c4e1e5e67607e6baadec9"></a>

<a id="canonical-38a090b48e3b76833f8617682331b86146060f8c7d34dba5ff4cbc7a124f17c6"></a>

## object_index property — items.system_metadata / e9f1122c1d84 / 10

Type: `"number"`. Computed.

Unique index for the object. Some objects need a unique integer index to be allocated for each
object type. This field will be populated for all objects that need it and will be zero otherwise.

- [owner_view](data-sources--site_registrations_by_site--reference--group-003.md#canonical-1d5660cd38b8c5bf3db286a39f459c4559b364b4b25fc230df01e99c93238da5): complete subsection reference.

<a id="canonical-342aabae4ca20174422a03e9f76af3c161cc86f3f8c5415cfb2f65386ebfa58b"></a>

<a id="canonical-9553b2e78c7d4d4dc8bc79298bfa0dbbcee91fa6dd736c30d11e65d9d13c6cdf"></a>

## tenant property — items.system_metadata / e9f1122c1d84 / 11

Type: `"string"`. Computed.

Tenant to which this configuration object belongs to. The value for this is found from presented
credentials.

<a id="canonical-0f61be7d3fb99e46538994e1050bf42401ef29b52d00cbfde5e9534b6120b692"></a>

<a id="canonical-8237e19d82c3b602d9ec882dc4d52fee1166ea51987f46ddc2dc10816efc32ac"></a>

## uid property — items.system_metadata / e9f1122c1d84 / 12

Type: `"string"`. Computed.

Uid is the unique in time and space value for this object. It is generated by the server on
successful creation of an object and is not allowed to change on Replace API. The value of is taken
from uid field of ObjectMetaType, if provided.

<a id="canonical-a925fa6bfae0146ebe34b291a81ddf0d8e087ca8732e436d98e4126cfd5be40a"></a>

## Next pages — items.system_metadata / e9f1122c1d84 / 13

- [items.system_metadata.initializers](data-sources--site_registrations_by_site--reference--group-003.md#canonical-d38e9f13f5a9efa85a0ea7e1a59fc2a4d1d976f4663ee4f28f090d4035bd744e)
- [items.system_metadata.labels](data-sources--site_registrations_by_site--reference--group-003.md#canonical-658639cc2e573c337edccd92785970404e32e104638d452e8245ec3e7b2193f3)
- [items.system_metadata.owner_view](data-sources--site_registrations_by_site--reference--group-003.md#canonical-1d5660cd38b8c5bf3db286a39f459c4559b364b4b25fc230df01e99c93238da5)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-d38e9f13f5a9efa85a0ea7e1a59fc2a4d1d976f4663ee4f28f090d4035bd744e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af4f0086aaacd911be4632ae01b37dc296e8b027dab488770f7d44e16c92a4e6"></a>

## items.system_metadata.initializers — items.system_metadata.initializers / 7eaad32dfe03 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.system_metadata](data-sources--site_registrations_by_site--reference--group-003.md#canonical-ad1a56f7a17b1bf18c3d85dc22c426ffab344a75a683a433efbf413a7d2b84fb)
- items.system_metadata.initializers

<a id="canonical-904e1903462244d3d55912e0d6b126742952757a626c9f9060aace09de1a53ab"></a>

Type: `"single"`. Computed.

Initializers tracks the progress of initialization of a configuration object.

<a id="canonical-db1d7acfc25c6ab68ce4fa581007b828f62de2c52a0de0981acdb5341decdfbe"></a>

## Direct properties — items.system_metadata.initializers / 7eaad32dfe03 / 3

- [pending](data-sources--site_registrations_by_site--reference--group-003.md#canonical-b850280977c2d2fedb2723f6ae85285fd44c74653da1027db175297b7a5f8d58): complete subsection reference.

- [result](data-sources--site_registrations_by_site--reference--group-003.md#canonical-a7d98eaabba9086c1392cc662dc4f17c55b3a994db76c5a1cb20f3a4c9d31adf): complete subsection reference.

<a id="canonical-932f9937c6306e6ce89ea138bd1b27fdd6d0f9345fb91602819edf9769f18aaa"></a>

## Next pages — items.system_metadata.initializers / 7eaad32dfe03 / 4

- [items.system_metadata.initializers.pending](data-sources--site_registrations_by_site--reference--group-003.md#canonical-b850280977c2d2fedb2723f6ae85285fd44c74653da1027db175297b7a5f8d58)
- [items.system_metadata.initializers.result](data-sources--site_registrations_by_site--reference--group-003.md#canonical-a7d98eaabba9086c1392cc662dc4f17c55b3a994db76c5a1cb20f3a4c9d31adf)
- [items.system_metadata](data-sources--site_registrations_by_site--reference--group-003.md#canonical-ad1a56f7a17b1bf18c3d85dc22c426ffab344a75a683a433efbf413a7d2b84fb)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-b850280977c2d2fedb2723f6ae85285fd44c74653da1027db175297b7a5f8d58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ce578521d9e18eb918c46c32e91efa705729547f28600832778985629d17638"></a>

## items.system_metadata.initializers.pending — items.system_metadata.initializers.pending / 2b6c65122efd / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.system_metadata](data-sources--site_registrations_by_site--reference--group-003.md#canonical-ad1a56f7a17b1bf18c3d85dc22c426ffab344a75a683a433efbf413a7d2b84fb)
- [items.system_metadata.initializers](data-sources--site_registrations_by_site--reference--group-003.md#canonical-d38e9f13f5a9efa85a0ea7e1a59fc2a4d1d976f4663ee4f28f090d4035bd744e)
- items.system_metadata.initializers.pending

<a id="canonical-82e26d5600ce1fc5586691e25bbd5e97f0188fcf481abb1c814bd99098a309c2"></a>

Type: `"list"`. Computed.

Pending is a list of initializers that must execute in order before this object is initialized. When
the last pending initializer is removed, and no failing result is set, the initializers struct will
be set to nil and the object is considered as initialized and visible to all clients.

<a id="canonical-86571c5cf5e00e966feb60ce4fcfd0ee989c7c517f79c70c38859c30f1741f7a"></a>

## Direct properties — items.system_metadata.initializers.pending / 2b6c65122efd / 3

<a id="canonical-655a92778fd543f813439831e7d010ee0989b022fd4f9c14c3cd271ed2e43090"></a>

<a id="canonical-97632f116a869f102c14da3b82da6fd97dff6de09d14b8ffdad97327c30bb68b"></a>

## name property — items.system_metadata.initializers.pending / 2b6c65122efd / 4

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

<a id="canonical-04304b55d316c4514c6283fa3e210a37b5a3b4766b35fc504493da3c75c2cfe6"></a>

## Next pages — items.system_metadata.initializers.pending / 2b6c65122efd / 5

- [items.system_metadata.initializers](data-sources--site_registrations_by_site--reference--group-003.md#canonical-d38e9f13f5a9efa85a0ea7e1a59fc2a4d1d976f4663ee4f28f090d4035bd744e)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-a7d98eaabba9086c1392cc662dc4f17c55b3a994db76c5a1cb20f3a4c9d31adf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4fa69b5bb09d796977745d424cea24fd95a63fefa356dd233bfa9692c495f6b"></a>

## items.system_metadata.initializers.result — items.system_metadata.initializers.result / 1b5f1071da47 / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.system_metadata](data-sources--site_registrations_by_site--reference--group-003.md#canonical-ad1a56f7a17b1bf18c3d85dc22c426ffab344a75a683a433efbf413a7d2b84fb)
- [items.system_metadata.initializers](data-sources--site_registrations_by_site--reference--group-003.md#canonical-d38e9f13f5a9efa85a0ea7e1a59fc2a4d1d976f4663ee4f28f090d4035bd744e)
- items.system_metadata.initializers.result

<a id="canonical-b6df2a7e4eff135d1626c3578fa656a9eff43a7427e7a032d61090a671e17f8b"></a>

Type: `"single"`. Computed.

Status is a return value for calls that don't return other objects.

<a id="canonical-7d113781b788244d668769f59be6c228c302103a06f36388104d3236b13a44b1"></a>

## Direct properties — items.system_metadata.initializers.result / 1b5f1071da47 / 3

<a id="canonical-b422e252e18d4f1034836e5066988619dc43499df36ef1c0e867fea12ef5e630"></a>

<a id="canonical-cab38a9f898cdb61115f43f955cd4eec92887e3dcb8cbfb0971837b2ba331b44"></a>

## code property — items.system_metadata.initializers.result / 1b5f1071da47 / 4

Type: `"number"`. Computed.

Suggested HTTP return code for this status, 0 if not set.

<a id="canonical-4f4ca29961c0ee34d7d1cce9cae3f455ef11f595177c5f6e0aaf89ad3432a373"></a>

<a id="canonical-ad7e8ea6056d384b60da8d619393e4a37309cd5ed7e9e38114880689edc09ba2"></a>

## reason property — items.system_metadata.initializers.result / 1b5f1071da47 / 5

Type: `"string"`. Computed.

Human-readable description of why this operation is in the 'Failure' status. If this value is empty
there is no information available.

<a id="canonical-6c4c39de572d5f213dc0bdbc18b9018e1145ce8c33509bcf95dd620a8e0e537f"></a>

<a id="canonical-be1d9c2022bcf571ddc5c88634b0a38ae35616b867ca60e3dd9d9dafa051b70a"></a>

## status property — items.system_metadata.initializers.result / 1b5f1071da47 / 6

Type: `"string"`. Computed.

Status of the operation. One of: 'Success' or 'Failure'.

<a id="canonical-0189a4d40e3d199be132b4a0a8a6093cbc4fb5f8da28fb98926beb3e86eaa0b4"></a>

## Next pages — items.system_metadata.initializers.result / 1b5f1071da47 / 7

- [items.system_metadata.initializers](data-sources--site_registrations_by_site--reference--group-003.md#canonical-d38e9f13f5a9efa85a0ea7e1a59fc2a4d1d976f4663ee4f28f090d4035bd744e)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-658639cc2e573c337edccd92785970404e32e104638d452e8245ec3e7b2193f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d986f9cad9a7fc8bc40cbc75b799de86ae4d72911149310fdaf84c8aeda2ae7"></a>

## items.system_metadata.labels — items.system_metadata.labels / f9db97ef375c / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.system_metadata](data-sources--site_registrations_by_site--reference--group-003.md#canonical-ad1a56f7a17b1bf18c3d85dc22c426ffab344a75a683a433efbf413a7d2b84fb)
- items.system_metadata.labels

<a id="canonical-f1980c541c45acbc439bb7df2fb9f4cbddce6ad116f16b2780e6ba64cf076237"></a>

Type: `"single"`. Computed.

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the operator or software. Values here can be interpreted by software(backend or
frontend) to enable certain behavior e.g. Things marked as soft-deleted(restorable).

<a id="canonical-d49576150826164e4ada4d28e552560e656db6de830807345b5165e8e9450665"></a>

## Direct properties — items.system_metadata.labels / f9db97ef375c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8793eb49edb3f6018d78b793a2174f26b0c0e82a43d2b82f773047655c882693"></a>

## Next pages — items.system_metadata.labels / f9db97ef375c / 4

- [items.system_metadata](data-sources--site_registrations_by_site--reference--group-003.md#canonical-ad1a56f7a17b1bf18c3d85dc22c426ffab344a75a683a433efbf413a7d2b84fb)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)

<a id="canonical-1d5660cd38b8c5bf3db286a39f459c4559b364b4b25fc230df01e99c93238da5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ae00ff83a762fd24ea16e014cd82160cedc0d47c6f8d6671611a4421cc7f57d"></a>

## items.system_metadata.owner_view — items.system_metadata.owner_view / ae0afd2ec61d / 2

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
- [Property reference](data-sources--site_registrations_by_site--reference--group-001.md#canonical-1248603df02ae3e61f32ee8b6fee3170e79477a7407bcfde66af4d7c062772c0)
- [items](data-sources--site_registrations_by_site--reference--group-001.md#canonical-230069c78b5f47cf8bc9f79e9a78d0c2ebdf9b6d5405a53c378ce64937a5863c)
- [items.system_metadata](data-sources--site_registrations_by_site--reference--group-003.md#canonical-ad1a56f7a17b1bf18c3d85dc22c426ffab344a75a683a433efbf413a7d2b84fb)
- items.system_metadata.owner_view

<a id="canonical-ac2a2e6318bf15e1a13ac938cf548b614a67efc36e3dc2e75defae1cf4dcd9d3"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

<a id="canonical-4008febb9e694d0225f4ab4ae7446964cfe2375bda9a36b6c21befae90b065b0"></a>

## Direct properties — items.system_metadata.owner_view / ae0afd2ec61d / 3

<a id="canonical-a790ddb993318b9b8afc60b957d2bcddec29aa2f96ecd328e2c936a96972b1c4"></a>

<a id="canonical-1dca9ed74d3916a0ced1c9521df243a55726324e94b914f71eed6ba99158ebbf"></a>

## kind property — items.system_metadata.owner_view / ae0afd2ec61d / 4

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="canonical-ae5647bac1425dd2c88b065e0317fd77ecae513fad9bd563cb5180f3fe481b1b"></a>

<a id="canonical-76975ae190af54a4683de12f8474c20e789de50ca7dafce0f086b6ed0d792743"></a>

## name property — items.system_metadata.owner_view / ae0afd2ec61d / 5

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

<a id="canonical-147ee17a276a09e134602134b8b75f8faa71688cfe0a2aec920f43db252e764f"></a>

<a id="canonical-93aa439305c71d5d31df1d2ae9240b76490dc95624a600273fcdfc741204fa55"></a>

## namespace property — items.system_metadata.owner_view / ae0afd2ec61d / 6

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

<a id="canonical-0596fd0ed7407192887867cb5f8f149857017a344fb224da9570d3ac47d23764"></a>

<a id="canonical-b988197599aa0a5a5c1c9756ea386e67f434b33f0b5878239600b213cfeccb87"></a>

## uid property — items.system_metadata.owner_view / ae0afd2ec61d / 7

Type: `"string"`. Computed.

UID. UID of the view object.

<a id="canonical-9d7db1ab0ea1758d0ebd07be5add9be3504e97ce5282268f7a52f7d53e4d5c01"></a>

## Next pages — items.system_metadata.owner_view / ae0afd2ec61d / 8

- [items.system_metadata](data-sources--site_registrations_by_site--reference--group-003.md#canonical-ad1a56f7a17b1bf18c3d85dc22c426ffab344a75a683a433efbf413a7d2b84fb)
- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-95cd454cd721919b77ea1bb44756386a23a0984b0ff7c5cf33cfc0a9a37682f8)
