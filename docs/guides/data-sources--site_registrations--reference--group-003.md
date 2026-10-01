---
page_title: "xcsh_site_registrations reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations reference."
---

# xcsh_site_registrations reference

<a id="canonical-0b4b5a684458435f1e684ae09d5530c980b315374a99006d5fe64e0fc0b58ce2"></a>

## items.object.system_metadata.initializers.pending — items.object.system_metadata.initializers.pending / cf793b542d29 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.object](data-sources--site_registrations--reference--group-002.md#canonical-eeb4fe90e1d3f233c434c0ab4bbdaa3daafeb37d22a781963dab076428488709)
- [items.object.system_metadata](data-sources--site_registrations--reference--group-002.md#canonical-026b416597acb8319a00e73a181ef399e84b9c0a61fb6caf2cc0cccbd00943a4)
- [items.object.system_metadata.initializers](data-sources--site_registrations--reference--group-002.md#canonical-a76e62137ccd64bbe9be83fa3eceda5e404a081b02044e054c341e68d9ba0490)
- items.object.system_metadata.initializers.pending

<a id="canonical-ec5ea6c884ef4434731b2024dbf0ac83488ad81bae63adc27dad04765e2b9741"></a>

Type: `"list"`. Computed.

Pending is a list of initializers that must execute in order before this object is initialized. When
the last pending initializer is removed, and no failing result is set, the initializers struct will
be set to nil and the object is considered as initialized and visible to all clients.

<a id="canonical-92bd4cc631f9586cbd089d4905b7036f9c5a98d3c7b9cb95364f1030387a1113"></a>

## Direct properties — items.object.system_metadata.initializers.pending / cf793b542d29 / 3

<a id="canonical-57c25e6ee5bd6c7b01e1569932b1427993a2e5c14a3892afea4ecbcc56e473f0"></a>

<a id="canonical-979166e22cba81228ba44c89507238b769540e3c0adbda6c42c6b678c1163c60"></a>

## name property — items.object.system_metadata.initializers.pending / cf793b542d29 / 4

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

<a id="canonical-175642e060ccb9f859af52169dd97265852d3796a5f9d156eae8acb6b5da9bc7"></a>

## Next pages — items.object.system_metadata.initializers.pending / cf793b542d29 / 5

- [items.object.system_metadata.initializers](data-sources--site_registrations--reference--group-002.md#canonical-a76e62137ccd64bbe9be83fa3eceda5e404a081b02044e054c341e68d9ba0490)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-c04be9f91d661f156657c2867a8a339dcef7c64764032f9dc230040f4f56771a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83707ef1ff88ed424ec60586c3fb3faa52357de6154647839c7eb967f0fb9c92"></a>

## items.object.system_metadata.initializers.result — items.object.system_metadata.initializers.result / 11ca4de62df8 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.object](data-sources--site_registrations--reference--group-002.md#canonical-eeb4fe90e1d3f233c434c0ab4bbdaa3daafeb37d22a781963dab076428488709)
- [items.object.system_metadata](data-sources--site_registrations--reference--group-002.md#canonical-026b416597acb8319a00e73a181ef399e84b9c0a61fb6caf2cc0cccbd00943a4)
- [items.object.system_metadata.initializers](data-sources--site_registrations--reference--group-002.md#canonical-a76e62137ccd64bbe9be83fa3eceda5e404a081b02044e054c341e68d9ba0490)
- items.object.system_metadata.initializers.result

<a id="canonical-94af63f7d0b17ed15fa20fde583667af91ffcc580f96dbd63bab06e512d128e2"></a>

Type: `"single"`. Computed.

Status is a return value for calls that don't return other objects.

<a id="canonical-b31be1a089ba5305a5078b928d287a011a7db66e35c170407f52e1c6b2d6674b"></a>

## Direct properties — items.object.system_metadata.initializers.result / 11ca4de62df8 / 3

<a id="canonical-c37921bfacc939f3c75def5ee8e47060ddc4f03c4d6fd387827f282ba4cc5f07"></a>

<a id="canonical-38302c901d9c43a226e484524d66c90dd13c1393c98d9070eb2e2e92d55e02b0"></a>

## code property — items.object.system_metadata.initializers.result / 11ca4de62df8 / 4

Type: `"number"`. Computed.

Suggested HTTP return code for this status, 0 if not set.

<a id="canonical-76ac0b12ce1b23503cf40024719fdaeff850cd3ec30ae1a5a931c476a527a015"></a>

<a id="canonical-c58c068440a755c4ec6ac7b041dcb5b67728363902af514b45bd541d1add19f4"></a>

## reason property — items.object.system_metadata.initializers.result / 11ca4de62df8 / 5

Type: `"string"`. Computed.

Human-readable description of why this operation is in the 'Failure' status. If this value is empty
there is no information available.

<a id="canonical-57347e1bf8fe119b268e6d57d363f000166b1ddc98f0a32bbae35ed02954521c"></a>

<a id="canonical-24106f728fe879cf978988d8be767b9950325aa51987eabd8e76cfcb475b2055"></a>

## status property — items.object.system_metadata.initializers.result / 11ca4de62df8 / 6

Type: `"string"`. Computed.

Status of the operation. One of: 'Success' or 'Failure'.

<a id="canonical-d88bbfdbbcd6e738850e2efc80bba7fedca454c170f26fc4c094f5ba3bce314c"></a>

## Next pages — items.object.system_metadata.initializers.result / 11ca4de62df8 / 7

- [items.object.system_metadata.initializers](data-sources--site_registrations--reference--group-002.md#canonical-a76e62137ccd64bbe9be83fa3eceda5e404a081b02044e054c341e68d9ba0490)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-75c1575b99fec8283c8506a3862ad067fe0806074f289bdddc96131706d89edb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87131f58ff1b69e067330cbec1115e5686272d6de575dfc47a9e37566538aaeb"></a>

## items.object.system_metadata.labels — items.object.system_metadata.labels / ca576cc88a11 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.object](data-sources--site_registrations--reference--group-002.md#canonical-eeb4fe90e1d3f233c434c0ab4bbdaa3daafeb37d22a781963dab076428488709)
- [items.object.system_metadata](data-sources--site_registrations--reference--group-002.md#canonical-026b416597acb8319a00e73a181ef399e84b9c0a61fb6caf2cc0cccbd00943a4)
- items.object.system_metadata.labels

<a id="canonical-14f5a382d300b33d405ccbeb206c987680436fdb3abc760b5c42b09765b5c363"></a>

Type: `"single"`. Computed.

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the operator or software. Values here can be interpreted by software(backend or
frontend) to enable certain behavior e.g. Things marked as soft-deleted(restorable).

<a id="canonical-ea099d42db1b48db5d04dfd174a3004e94821c7b6a3ecfc2b5c94c44bde439c6"></a>

## Direct properties — items.object.system_metadata.labels / ca576cc88a11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb01ed756968facb05add3e99c8166d518cdfa18e59bb228ea71280b839b3fe4"></a>

## Next pages — items.object.system_metadata.labels / ca576cc88a11 / 4

- [items.object.system_metadata](data-sources--site_registrations--reference--group-002.md#canonical-026b416597acb8319a00e73a181ef399e84b9c0a61fb6caf2cc0cccbd00943a4)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-90bfbfab27183d7dc0c13bd27364bc18989be4cf49c55df542e1639b47dfd778"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d314305d2003e79f2e10d2748a6f6977d1a187e2b560350a11cf7e26e6817ba2"></a>

## items.object.system_metadata.namespace — items.object.system_metadata.namespace / 5fd2c61e2e55 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.object](data-sources--site_registrations--reference--group-002.md#canonical-eeb4fe90e1d3f233c434c0ab4bbdaa3daafeb37d22a781963dab076428488709)
- [items.object.system_metadata](data-sources--site_registrations--reference--group-002.md#canonical-026b416597acb8319a00e73a181ef399e84b9c0a61fb6caf2cc0cccbd00943a4)
- items.object.system_metadata.namespace

<a id="canonical-a17081f02131d8119215e5bbbf174932a30c13f789a47154661a104795899c7d"></a>

Type: `"list"`. Computed.

The namespace this object belongs to. This is populated by the service based on the
metadata.namespace field when an object is created.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

<a id="canonical-e4267371fa1c85a63de047d21a0bc8a2b51c2ecb9789e26320a6acb6a637dcb4"></a>

## Direct properties — items.object.system_metadata.namespace / 5fd2c61e2e55 / 3

<a id="canonical-86b1d4b76838a390234bc81441c9846a353ae74bccfcf558eec71958ea7d7a45"></a>

<a id="canonical-23621ee1f5497a945d5b5aa7e6e0fb8952237a213a18ddf046e0fb128e78d2d7"></a>

## kind property — items.object.system_metadata.namespace / 5fd2c61e2e55 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

<a id="canonical-9c7875f1ba428687bd2d67891712a88287ca4cf9ebaf92ac37b047bcd905fbbc"></a>

<a id="canonical-9bfe42fe4106d8210fa9d357e2d6ea0a2671301bc758ff3365f243e1aa2780dd"></a>

## name property — items.object.system_metadata.namespace / 5fd2c61e2e55 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

<a id="canonical-906ea59cc7f9d0dc3c68875a21d76e411dc9a1b74454456826b796b349ed187e"></a>

<a id="canonical-1ddd962109f41bc73865ed32cf54af0850df72d526832a14872e760704d37032"></a>

## namespace property — items.object.system_metadata.namespace / 5fd2c61e2e55 / 6

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

<a id="canonical-2f3fe1ef119438841dd3d9c2b3a1751abcd0ced0be3a7d9c110e73fbde795c2c"></a>

<a id="canonical-58dbf7f3d27d97c7d1a293f2048851851b5ccd4b538815c72774e6b09c59da21"></a>

## tenant property — items.object.system_metadata.namespace / 5fd2c61e2e55 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

<a id="canonical-83e920295646dd23a2ec6b0366ed09c17635e6bfc47b71b89e5ee870ab164d06"></a>

<a id="canonical-8d4efc222767ae4fdac6603fc4739e3009769ee7920bb79eff6a234ef4630a93"></a>

## uid property — items.object.system_metadata.namespace / 5fd2c61e2e55 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

<a id="canonical-39b20ba3f933c7465e9b8e819440d41770825f225157103c6ab9d333cdac836c"></a>

## Next pages — items.object.system_metadata.namespace / 5fd2c61e2e55 / 9

- [items.object.system_metadata](data-sources--site_registrations--reference--group-002.md#canonical-026b416597acb8319a00e73a181ef399e84b9c0a61fb6caf2cc0cccbd00943a4)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-ea691647d54d50006667f246619c14e941822dfc02119518acb77914d4f058d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef20488d6faa1e13cb370f2d2e38a704e15b8324a5d255d56fc202ad6e0c68a1"></a>

## items.object.system_metadata.owner_view — items.object.system_metadata.owner_view / 0600c24dd3cd / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.object](data-sources--site_registrations--reference--group-002.md#canonical-eeb4fe90e1d3f233c434c0ab4bbdaa3daafeb37d22a781963dab076428488709)
- [items.object.system_metadata](data-sources--site_registrations--reference--group-002.md#canonical-026b416597acb8319a00e73a181ef399e84b9c0a61fb6caf2cc0cccbd00943a4)
- items.object.system_metadata.owner_view

<a id="canonical-049e4e88cc9daa29737b8b11a4f43c298ead88e475cc7df782f77fdc77edb866"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

<a id="canonical-acdd9f77290662c144cdd423c68919613396aca1a74df049b7d58023e4961275"></a>

## Direct properties — items.object.system_metadata.owner_view / 0600c24dd3cd / 3

<a id="canonical-c1413c85836d212b65806d9577e098b21bd5f618a22c0389310c41820e690e24"></a>

<a id="canonical-2377d7a0b4da9898d82c0bd35ac471f38e79f5192e5a8551354edc0cd9d6c15a"></a>

## kind property — items.object.system_metadata.owner_view / 0600c24dd3cd / 4

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="canonical-8d8e2fcddc9d7ea7aed8e0b0ddd28e7d9ab1c335ee29f5ee95efda4d5080aca8"></a>

<a id="canonical-b8a0851e858116e5b5ac3b6e10cb05cef6eb31363e811cf9dc66c36b67d1c266"></a>

## name property — items.object.system_metadata.owner_view / 0600c24dd3cd / 5

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

<a id="canonical-93d16a76a2636983df375964dd5d6e2810e6f420e5c5b42f235a8ecc54dc2f79"></a>

<a id="canonical-88aee71c6e2674bce629b2232b602f56f568afb62d5120a2114c4198b0305652"></a>

## namespace property — items.object.system_metadata.owner_view / 0600c24dd3cd / 6

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

<a id="canonical-13df77cc10f432efc9ad49420bca883c924ee7effbbc14a6da9f96f3bae71b4b"></a>

<a id="canonical-198fb65e278f7562fcb24d96a97677eb3841d3015809eccbc25e2fe93f593f44"></a>

## uid property — items.object.system_metadata.owner_view / 0600c24dd3cd / 7

Type: `"string"`. Computed.

UID. UID of the view object.

<a id="canonical-3c979f77c2f31e45fd1833e6e0ff3d35bb4132023cd7c886a83af63c2e68c9ff"></a>

## Next pages — items.object.system_metadata.owner_view / 0600c24dd3cd / 8

- [items.object.system_metadata](data-sources--site_registrations--reference--group-002.md#canonical-026b416597acb8319a00e73a181ef399e84b9c0a61fb6caf2cc0cccbd00943a4)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-7ac394f34f8d5c276a318621b114222aca6843de080848e27f159a632a7c176f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc6e4cbdd09581950a6625e45abebaa2ab2f30b4ca1671c0db4f359a14aa7a20"></a>

## items.owner_view — items.owner_view / 3015d33752c2 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- items.owner_view

<a id="canonical-233ccbd925f0a8af37f03c16fd1fd023e64b27557bf93cccf41aef591b2b4518"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

<a id="canonical-2560fa2a1258d856777c6addf714429dc837051c6ffb8eb58d9ab16f55d943bd"></a>

## Direct properties — items.owner_view / 3015d33752c2 / 3

<a id="canonical-2e23e7f15a975c7c39bd67fdefd823de6df2fec7f4a3f8c329ae43fc422c9d11"></a>

<a id="canonical-f03a22789566eb839602b4c17d246f3cfcb128fa265c95425e58d2ea74df1827"></a>

## kind property — items.owner_view / 3015d33752c2 / 4

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="canonical-5b69c7627373520a16df90740cbb6268d541eca9579166289036a01b20f526ed"></a>

<a id="canonical-98af066856b8abed7b9d0b76061c4316f164efd2590acd9727d2cdd147b13d17"></a>

## name property — items.owner_view / 3015d33752c2 / 5

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

<a id="canonical-4d9655866bea4f54e2ed34d54c48be47380c8ea13c5c3f2aa7bf3c892c21f20b"></a>

<a id="canonical-aaec5b1e2c9076337ac8e92beb04424aa6bf57ddb62c4356935f52944b443112"></a>

## namespace property — items.owner_view / 3015d33752c2 / 6

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

<a id="canonical-591505a07838ece9a8cac208c915734267990a94619bc1b1f7a15844e9762bb3"></a>

<a id="canonical-73dd80f7a6a51356c54cb7630a68f894e0bef8920357f3a1ed032a252d5c9097"></a>

## uid property — items.owner_view / 3015d33752c2 / 7

Type: `"string"`. Computed.

UID. UID of the view object.

<a id="canonical-1af606cfc937b1031f50ae94eb1b21f4b911f5aaa76f1766a7926f3b825acca8"></a>

## Next pages — items.owner_view / 3015d33752c2 / 8

- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-3a0368175e24ddfae32c24a9f0f1bf9115db0bfcd0659025cc35d1e4ceec8a63"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1643a27f7138d8ca8959a46171c25b08efab171cf929d49d1c2e8426ea0e272a"></a>

## items.system_metadata — items.system_metadata / 8bede6c138d6 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- items.system_metadata

<a id="canonical-86192061a67d13bc67b8f3dcbbe18f331e38daab87e49d994d6a62c6e64e6f80"></a>

Type: `"single"`. Computed.

SystemObjectGetMetaType is metadata generated or populated by the system for all persisted objects
and cannot be updated directly by users.

<a id="canonical-efe5bd4b261aefc9cffdfebdbf9347eec7d46dc51da5f7635a5376653f84221c"></a>

## Direct properties — items.system_metadata / 8bede6c138d6 / 3

<a id="canonical-efd9310f93e64891212adecd4187522c165ae8cc3c16dca616f92e0446976478"></a>

<a id="canonical-d06ec65c6d347c23803cbed10350dea98713d61cac7cb6b9dd554e58defacfe6"></a>

## creation_timestamp property — items.system_metadata / 8bede6c138d6 / 4

Type: `"string"`. Computed.

CreationTimestamp is a timestamp representing the server time when this object was created. It is
not guaranteed to be set in happens-before order across separate operations. Clients may not set
this value.

<a id="canonical-4b26f0de5dfecdc8bd518a402747cde9c93cff298bebdfc2400ba62c3df7c46f"></a>

<a id="canonical-e5ed85df9deb53c370593cd5e4b15021f34ac1ea43391414c5e9b4dfbce2d3bc"></a>

## creator_class property — items.system_metadata / 8bede6c138d6 / 5

Type: `"string"`. Computed.

Value identifying the class of the user or service which created this configuration object.

<a id="canonical-f16190d36f1c3f71fb227db5a6d73fba0fa8fe94b3cf9c98e5f31671ceebf860"></a>

<a id="canonical-e95c58a8a7fbb89c2788b80354e1f2c1d8f39dcf80edcfd1be66e79efa495005"></a>

## creator_id property — items.system_metadata / 8bede6c138d6 / 6

Type: `"string"`. Computed.

Value identifying the exact user or service that created this configuration object.

<a id="canonical-7134e98e7d5bd5cf5f37475819ad648705df5317136ef37033c84c4445ebd225"></a>

<a id="canonical-1d672ea42db4f831973404738972c57a991b7e0d9f77cfa9264fdce5d4f0e43f"></a>

## deletion_timestamp property — items.system_metadata / 8bede6c138d6 / 7

Type: `"string"`. Computed.

DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is
set by the server when a graceful deletion is requested by the user, and is not directly settable by
a client. The resource is expected to be deleted (no longer visible from resource lists, and not..

<a id="canonical-4f07f64fcd0506351735359160bf2eafb465c3faea68b698ef34294f5769a991"></a>

<a id="canonical-e7bb5427ccc0a9a1895e8b33ed1713d3a734506489f58d1592c02accf4b44bba"></a>

## finalizers property — items.system_metadata / 8bede6c138d6 / 8

Type: `["list", "string"]`. Computed.

Must be empty before the object is deleted from the registry. Each entry is an identifier for the
responsible component that will remove the entry from the list. If the deletionTimestamp of the
object is non-nil, entries in this list can only be removed.

- [initializers](data-sources--site_registrations--reference--group-003.md#canonical-b0f4e994a594cab1a92b33aba7944448904c325e88bab1276725e95aebfb8727): complete subsection reference.

- [labels](data-sources--site_registrations--reference--group-003.md#canonical-3384dd4fd65b5335c5bb592b69051c2a84c67dba366324d779ff839ee45035e0): complete subsection reference.

<a id="canonical-2403c97d577c14b405254864909eb1225fd21c32228f1b5e7ed73a251d8781bd"></a>

<a id="canonical-2da9b4fa9355eea56a6b3edff5885748de866e92e2989a41139ff15dd28f1fe6"></a>

## modification_timestamp property — items.system_metadata / 8bede6c138d6 / 9

Type: `"string"`. Computed.

ModificationTimestamp is a timestamp representing the server time when this object was last
modified.

<a id="canonical-74f3c17553320a4958af3aed9444007f98a23814d4c708b7d3bef9e7b5a702fd"></a>

<a id="canonical-1659f695d36736a9ebd46e7534d69e683f7875d1fe10c511fdbd391c8e21fdbe"></a>

## object_index property — items.system_metadata / 8bede6c138d6 / 10

Type: `"number"`. Computed.

Unique index for the object. Some objects need a unique integer index to be allocated for each
object type. This field will be populated for all objects that need it and will be zero otherwise.

- [owner_view](data-sources--site_registrations--reference--group-003.md#canonical-d520399cd493c43e2c6ef11da2f3c7837c00177df396e4cf13d717a0b73cfeed): complete subsection reference.

<a id="canonical-cadec740b4772bb87dc7a0ead2d184f05e786f4574044f7252166d049ca26f8a"></a>

<a id="canonical-20221bd5b947aaf3bbbcd445a6c769d27b26e46842fa294c4486f0dfec0532ba"></a>

## tenant property — items.system_metadata / 8bede6c138d6 / 11

Type: `"string"`. Computed.

Tenant to which this configuration object belongs to. The value for this is found from presented
credentials.

<a id="canonical-b38d77e478cf6c3eaf3ea5267ca3d077d393da82bfc5cdf08882d45e5082aa9c"></a>

<a id="canonical-dcbe74d71025a4257e8c1bfda6aba3ace801fe0727bd6659668f2e1c915b5a93"></a>

## uid property — items.system_metadata / 8bede6c138d6 / 12

Type: `"string"`. Computed.

Uid is the unique in time and space value for this object. It is generated by the server on
successful creation of an object and is not allowed to change on Replace API. The value of is taken
from uid field of ObjectMetaType, if provided.

<a id="canonical-0f706457a3346831d2d90a3c44e59478f838c17af87bf6cd6312f834e3e0f471"></a>

## Next pages — items.system_metadata / 8bede6c138d6 / 13

- [items.system_metadata.initializers](data-sources--site_registrations--reference--group-003.md#canonical-b0f4e994a594cab1a92b33aba7944448904c325e88bab1276725e95aebfb8727)
- [items.system_metadata.labels](data-sources--site_registrations--reference--group-003.md#canonical-3384dd4fd65b5335c5bb592b69051c2a84c67dba366324d779ff839ee45035e0)
- [items.system_metadata.owner_view](data-sources--site_registrations--reference--group-003.md#canonical-d520399cd493c43e2c6ef11da2f3c7837c00177df396e4cf13d717a0b73cfeed)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-b0f4e994a594cab1a92b33aba7944448904c325e88bab1276725e95aebfb8727"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bece7f8b3a435da682bd70c29c2054ac62f38336953d7fad2ee1e1fe6a271a2"></a>

## items.system_metadata.initializers — items.system_metadata.initializers / c82ecbd6e6c0 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.system_metadata](data-sources--site_registrations--reference--group-003.md#canonical-3a0368175e24ddfae32c24a9f0f1bf9115db0bfcd0659025cc35d1e4ceec8a63)
- items.system_metadata.initializers

<a id="canonical-63d3ca7329ebe8a0b4b627620e65c652dc1ffedf99a621e0cc81085083f3cc67"></a>

Type: `"single"`. Computed.

Initializers tracks the progress of initialization of a configuration object.

<a id="canonical-e39906c08faf0869dc2f57ed64127e5fd2572c35ad6c1f5aec491076d8c9e1f7"></a>

## Direct properties — items.system_metadata.initializers / c82ecbd6e6c0 / 3

- [pending](data-sources--site_registrations--reference--group-003.md#canonical-5db468e58075c5056ec5eaabd87b2cf49d29fafc9062ba8fd99fe7a8d1a69cf6): complete subsection reference.

- [result](data-sources--site_registrations--reference--group-003.md#canonical-169821d3575328901c421fa1163682d28f96dd0137fe6f99037783dd665858b9): complete subsection reference.

<a id="canonical-eb276c18069e6b46bcdf29e0234cd538f435f1ead32cd5f26336c75d3aadbc0e"></a>

## Next pages — items.system_metadata.initializers / c82ecbd6e6c0 / 4

- [items.system_metadata.initializers.pending](data-sources--site_registrations--reference--group-003.md#canonical-5db468e58075c5056ec5eaabd87b2cf49d29fafc9062ba8fd99fe7a8d1a69cf6)
- [items.system_metadata.initializers.result](data-sources--site_registrations--reference--group-003.md#canonical-169821d3575328901c421fa1163682d28f96dd0137fe6f99037783dd665858b9)
- [items.system_metadata](data-sources--site_registrations--reference--group-003.md#canonical-3a0368175e24ddfae32c24a9f0f1bf9115db0bfcd0659025cc35d1e4ceec8a63)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-5db468e58075c5056ec5eaabd87b2cf49d29fafc9062ba8fd99fe7a8d1a69cf6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e8ae6914201538e5f7e9430522fd6dd7bdd097056aa1a2e5a301026807d470e"></a>

## items.system_metadata.initializers.pending — items.system_metadata.initializers.pending / 043e3b181fbf / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.system_metadata](data-sources--site_registrations--reference--group-003.md#canonical-3a0368175e24ddfae32c24a9f0f1bf9115db0bfcd0659025cc35d1e4ceec8a63)
- [items.system_metadata.initializers](data-sources--site_registrations--reference--group-003.md#canonical-b0f4e994a594cab1a92b33aba7944448904c325e88bab1276725e95aebfb8727)
- items.system_metadata.initializers.pending

<a id="canonical-02ca139dac8813bfb5975f96c650afc6aded5d36d7cc0c78d4f5f8f1483089ae"></a>

Type: `"list"`. Computed.

Pending is a list of initializers that must execute in order before this object is initialized. When
the last pending initializer is removed, and no failing result is set, the initializers struct will
be set to nil and the object is considered as initialized and visible to all clients.

<a id="canonical-2c9ff2193f9ff6ba104937cd18a3545e443a4bac12281d681a5b11b3e6ba0326"></a>

## Direct properties — items.system_metadata.initializers.pending / 043e3b181fbf / 3

<a id="canonical-c41bc47e4e5f14e03350ebec39e5d8f37e4358fc10c04c4f7fd15cc90ed2e796"></a>

<a id="canonical-6e9eb09c0716208ed9c6b1132f162c76e36bd83ebfecc5b2481bee67fdbf6f94"></a>

## name property — items.system_metadata.initializers.pending / 043e3b181fbf / 4

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

<a id="canonical-f5925a3b91983d3149e15ebe349142b89590eda2e540b9c623eb2521fb3b32bd"></a>

## Next pages — items.system_metadata.initializers.pending / 043e3b181fbf / 5

- [items.system_metadata.initializers](data-sources--site_registrations--reference--group-003.md#canonical-b0f4e994a594cab1a92b33aba7944448904c325e88bab1276725e95aebfb8727)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-169821d3575328901c421fa1163682d28f96dd0137fe6f99037783dd665858b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36a5d62caf8105bd4b26a0245b5b20b4fe63007e7a98ecf66f813016ea4609d7"></a>

## items.system_metadata.initializers.result — items.system_metadata.initializers.result / 5b04493fdb03 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.system_metadata](data-sources--site_registrations--reference--group-003.md#canonical-3a0368175e24ddfae32c24a9f0f1bf9115db0bfcd0659025cc35d1e4ceec8a63)
- [items.system_metadata.initializers](data-sources--site_registrations--reference--group-003.md#canonical-b0f4e994a594cab1a92b33aba7944448904c325e88bab1276725e95aebfb8727)
- items.system_metadata.initializers.result

<a id="canonical-d4c93cacc9b7fb78eeee4e6f8af7afef48b4cdba6f295c55e87c1f0b069c7349"></a>

Type: `"single"`. Computed.

Status is a return value for calls that don't return other objects.

<a id="canonical-ead23ffb9fee9995957afc02c27e30b23ef898e981950ad63db6c345ce5fe6d1"></a>

## Direct properties — items.system_metadata.initializers.result / 5b04493fdb03 / 3

<a id="canonical-72ecc04e80b2944f105feecfd309a9f43f24363e14f96283438f2e8df665ea24"></a>

<a id="canonical-5fc7590c2c60b5d828f65a0c13b9fbbf8b5097ffd81e8a371389a39b53094fb9"></a>

## code property — items.system_metadata.initializers.result / 5b04493fdb03 / 4

Type: `"number"`. Computed.

Suggested HTTP return code for this status, 0 if not set.

<a id="canonical-08a0b18eba662973aa1b3f12b4205e6a3fd4ca53a21d684ec5db25b8f4c280d7"></a>

<a id="canonical-ad0be5c164d9cf8bd7d0b7bd43c6dc830796c5797e5c291eee4be911ac038d5a"></a>

## reason property — items.system_metadata.initializers.result / 5b04493fdb03 / 5

Type: `"string"`. Computed.

Human-readable description of why this operation is in the 'Failure' status. If this value is empty
there is no information available.

<a id="canonical-d20d96ea8ec7b768a6a16335a56f19951ea60fc4d4e6c4ab9624fd8751f75ecd"></a>

<a id="canonical-0b242741865ab312fcdc54efb7e20a4bc8fe05ae43d5f483b546c822eda43ef0"></a>

## status property — items.system_metadata.initializers.result / 5b04493fdb03 / 6

Type: `"string"`. Computed.

Status of the operation. One of: 'Success' or 'Failure'.

<a id="canonical-8e8f016a77fbc88f4430a430c47f7154db6d2df9d2008610acfdde23d7b1986c"></a>

## Next pages — items.system_metadata.initializers.result / 5b04493fdb03 / 7

- [items.system_metadata.initializers](data-sources--site_registrations--reference--group-003.md#canonical-b0f4e994a594cab1a92b33aba7944448904c325e88bab1276725e95aebfb8727)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-3384dd4fd65b5335c5bb592b69051c2a84c67dba366324d779ff839ee45035e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f7b0f0c85a44cd9695764d3f1bca77eed7b3fbdc2c2d9d3d16f60b87088ba70"></a>

## items.system_metadata.labels — items.system_metadata.labels / 2705ceb1459b / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.system_metadata](data-sources--site_registrations--reference--group-003.md#canonical-3a0368175e24ddfae32c24a9f0f1bf9115db0bfcd0659025cc35d1e4ceec8a63)
- items.system_metadata.labels

<a id="canonical-a0e3f76b042082e53510bca2e80cdfb23f605494f287a479cdb9cfe625b24568"></a>

Type: `"single"`. Computed.

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the operator or software. Values here can be interpreted by software(backend or
frontend) to enable certain behavior e.g. Things marked as soft-deleted(restorable).

<a id="canonical-68552957cd92479483a54fb3664398b499871d59dc822ca2894e424587dc9d32"></a>

## Direct properties — items.system_metadata.labels / 2705ceb1459b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-687d3ac61fea7ee6f4de802a7b5e18c06b22f64caae18c7150106dc4aef9fbf1"></a>

## Next pages — items.system_metadata.labels / 2705ceb1459b / 4

- [items.system_metadata](data-sources--site_registrations--reference--group-003.md#canonical-3a0368175e24ddfae32c24a9f0f1bf9115db0bfcd0659025cc35d1e4ceec8a63)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)

<a id="canonical-d520399cd493c43e2c6ef11da2f3c7837c00177df396e4cf13d717a0b73cfeed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3af7ba1cc128da27f34e1ba99c4d5b15085eb51a2b40dca24bbcbd34e690221"></a>

## items.system_metadata.owner_view — items.system_metadata.owner_view / 4fb823217581 / 2

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
- [Property reference](data-sources--site_registrations--reference--group-001.md#canonical-e9ce9f37e9d5fb7e5425df57d652e74015a83817a0e9796c5622ad59128f54cc)
- [items](data-sources--site_registrations--reference--group-001.md#canonical-97c8117489f9715e96dfb0465e0afb6d1c4108a03f1daf42d03868718f1b2b7e)
- [items.system_metadata](data-sources--site_registrations--reference--group-003.md#canonical-3a0368175e24ddfae32c24a9f0f1bf9115db0bfcd0659025cc35d1e4ceec8a63)
- items.system_metadata.owner_view

<a id="canonical-61397b4e811140813f6f3ae5158e9d65dc85b7bd8d550fa975c122418731e86e"></a>

Type: `"single"`. Computed.

ViewRefType represents a reference to a view.

<a id="canonical-71e95b0ca100690aeed8270e3ef3222517a5c0e3d7881d638a885dcbfa7f2d9f"></a>

## Direct properties — items.system_metadata.owner_view / 4fb823217581 / 3

<a id="canonical-6da1836d496e175a6f9030039e07cbefafb4ab30366ad2d9f5dcd13710a19bb9"></a>

<a id="canonical-47ae6c8f188480ff26eb5a5e2dd0f608bb75cc104ebf941a446dbbd1b9324037"></a>

## kind property — items.system_metadata.owner_view / 4fb823217581 / 4

Type: `"string"`. Computed.

Kind. Kind of the view object.

<a id="canonical-4bfc113a3af61837f09be162fb0d236b43e589d356b8d62a81de19468faf2589"></a>

<a id="canonical-86b5e2d777c754d245836a7b5cf68cbc85817ebab8b11cd4a463add5a8c10d8b"></a>

## name property — items.system_metadata.owner_view / 4fb823217581 / 5

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

<a id="canonical-2f0be0424b8524abfe04dbb2e6ebda1d3afb01e346453fc45806b63341398aed"></a>

<a id="canonical-c2c615a001ea324017ab00fa2d97652836635ebb71f854ca33384d726bb59479"></a>

## namespace property — items.system_metadata.owner_view / 4fb823217581 / 6

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

<a id="canonical-571e77e312acad2171ef4f8a0daf4f36b73f7682fc6c87f172970516347ea70e"></a>

<a id="canonical-8b806ab0be18fbadd0e4986336a36563ab97b8e39c7510b8ee2f172ab9396810"></a>

## uid property — items.system_metadata.owner_view / 4fb823217581 / 7

Type: `"string"`. Computed.

UID. UID of the view object.

<a id="canonical-5f7c78966030652e441048953c4654bfed7651bc677939ee98758f7f9b21a786"></a>

## Next pages — items.system_metadata.owner_view / 4fb823217581 / 8

- [items.system_metadata](data-sources--site_registrations--reference--group-003.md#canonical-3a0368175e24ddfae32c24a9f0f1bf9115db0bfcd0659025cc35d1e4ceec8a63)
- [xcsh_site_registrations](../data-sources/site_registrations.md#canonical-b81461a3656491bbffcae10d2e13be349aedf1a9b4d7c19fee793bf65a7c1323)
