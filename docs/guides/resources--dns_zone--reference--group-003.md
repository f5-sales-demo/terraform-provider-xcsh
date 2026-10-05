---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-1230200010210003-2223301210223102-1103002303320013-2120023000212111-0321233110112030-2313230233232022-1200231211230102-3022001212231321"></a>

## primary.rr_set_group.rr_set.cds_record — cds_record / 320300332131 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.cds_record

<a id="canonical-3012220321222313-2102100303223132-3320001303231312-1113220202122233-3233033202202311-3303302130320322-3130001301210120-3230110112300123"></a>

Type: `"object"`. single nested block, Optional.

DNS CDS Record. DNS CDS Record.

Upstream description:

DNS CDS Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
cds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110121023220131-0112103103111233-1103110030110032-2113210103113233-0232130033131322-2310320112313032-0033313120130032-3313000322222031"></a>

## Direct properties — cds_record / 320300332131 / 3

<a id="canonical-1202020213032001-1221001000120313-0200233210221010-1120001323311232-0231212323131302-1220021102102332-1111103201120221-0023103032201132"></a>

<a id="canonical-0122321102113013-1123212013311321-3221302311033230-3303303012322231-0221320012122012-2231301311000023-2130221300233202-1231012330021311"></a>

## name property — cds_record / 320300332131 / 4

Type: `"string"`. Optional.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033): complete subsection reference.

<a id="canonical-2021121311011300-2131331220222120-1122213023102211-2333211030203321-2223003223000032-0232121200330020-3130020322301300-3330233330002101"></a>

## Next pages — cds_record / 320300332131 / 5

- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-003.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001300212220323-1333231020303312-0212021133010032-0022311233222100-2002222030332102-3303013031010311-3223231310113220-2121113133112121"></a>

## primary.rr_set_group.rr_set.cds_record.values — values / 012312233223 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- primary.rr_set_group.rr_set.cds_record.values

<a id="canonical-2321310003223320-1321332333030003-2322033021331223-1330322002030000-3221210101203310-1200302130113113-0210313223323033-2023023212230202"></a>

Type: `"object"`. list nested block, Optional.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key_tag"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha256_digest"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha384_digest"),
  validators.ConflictingListObjectAttributes("sha256_digest",
    "sha384_digest")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020312121300020-0030303011322313-3233123330123121-3001310330032321-0021212321223122-1022301001111311-3021200232111231-3312123320300330"></a>

## Direct properties — values / 012312233223 / 3

<a id="canonical-3221313023302011-1302331132022102-2102313232003220-1011111302030030-2132232010020200-3212131221320011-0101312021001313-2011310131011331"></a>

<a id="canonical-1112310213322130-3011110303232003-1001211032030213-3023212223000330-1311102030332031-1011031033223112-2231031032002012-0203233203300101"></a>

## ds_key_algorithm property — values / 012312233223 / 4

Type: `"string"`. Optional.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key-value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ECDSAP256SHA256","ECDSAP384SHA384","ED25519","ED448","RSASHA1","RSASHA1NSEC3SHA1","RSASHA256","RSASHA512","UNSPECIFIED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3231012122013300-0201102022233213-3230031002000002-1022233323000132-1132223033013323-1210102310333003-0201113120223330-1102111303211100"></a>

<a id="canonical-0121032023333122-2312113300330302-2022120123200113-0000012021313102-1200231120122021-0201013202223112-0302003013220113-1010000311333201"></a>

## key_tag property — values / 012312233223 / 5

Type: `"number"`. Optional.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](resources--dns_zone--reference--group-003.md#canonical-3100120201012212-2003022220310300-3201231103212000-2113003003200130-0200023032301110-3232330220223333-2033203112000101-0103201010021303): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-003.md#canonical-2302310320330210-0313333002331000-2112300301232232-3322031231010011-2101010213322220-3111320011201223-3300032310020120-3321323123001031): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-003.md#canonical-2032220222001301-1021003012000102-0101002330300323-0312203302002020-0003200112112211-3021032200112002-2003312333133200-3113331200110221): complete subsection reference.

<a id="canonical-3132300013023032-3013132331211011-3030333033333330-0020230302302302-1221002231310123-0203232222210311-1301013112101132-0020031321323321"></a>

## Next pages — values / 012312233223 / 6

- [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](resources--dns_zone--reference--group-003.md#canonical-3100120201012212-2003022220310300-3201231103212000-2113003003200130-0200023032301110-3232330220223333-2033203112000101-0103201010021303)
- [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](resources--dns_zone--reference--group-003.md#canonical-2302310320330210-0313333002331000-2112300301232232-3322031231010011-2101010213322220-3111320011201223-3300032310020120-3321323123001031)
- [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](resources--dns_zone--reference--group-003.md#canonical-2032220222001301-1021003012000102-0101002330300323-0312203302002020-0003200112112211-3021032200112002-2003312333133200-3113331200110221)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3100120201012212-2003022220310300-3201231103212000-2113003003200130-0200023032301110-3232330220223333-2033203112000101-0103201010021303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231033010133020-1322210223030022-2130020213030200-2202133323312322-3232300103132021-3211202010311000-2232322123231103-2300233303002311"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha1_digest — sha1_digest / 122023003223 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-003.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- primary.rr_set_group.rr_set.cds_record.values.sha1_digest

<a id="canonical-3033132103122202-2320322021002103-2212203030021003-3313002010003333-0320010103312010-3332021022111200-0130221032312321-1230113302010031"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
```

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

Terraform syntax:

```terraform
sha1_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330210213200211-1002100333321323-2103321323100132-3332010020101023-0001210012021130-0230113211321122-2120023310301332-2322011223113222"></a>

## Direct properties — sha1_digest / 122023003223 / 3

<a id="canonical-2132020012111003-1303102023010130-1010202313311110-2100201300203203-2201121330113311-0220010203302022-2223132130313211-3312311030322213"></a>

<a id="canonical-0331112220113100-3101100001222013-1012301333220330-3022231112121031-0202001122101233-2330322123013321-3330231303300010-1033021313122030"></a>

## digest property — sha1_digest / 122023003223 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-0232133020320220-1233010330110302-0311131011020223-1231123203100033-2232321232212332-3300332303200022-1223011012322131-0131300213333332"></a>

## Next pages — sha1_digest / 122023003223 / 5

- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-003.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2302310320330210-0313333002331000-2112300301232232-3322031231010011-2101010213322220-3111320011201223-3300032310020120-3321323123001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202331012030300-1203311110323110-2213310203233112-1102301301032023-3102313310110112-1012222012023000-1102213130102322-2222301122132302"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha256_digest — sha256_digest / 302020022020 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-003.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- primary.rr_set_group.rr_set.cds_record.values.sha256_digest

<a id="canonical-0010233123223233-2012102031300132-3312323022113110-3031311102102121-2101311323203202-2131032213233033-1131311220010212-2221222002001002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
```

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

Terraform syntax:

```terraform
sha256_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122112010001212-0100110021312203-1320013310131210-2003013121000023-2101200323323223-0220021113303213-0220323222330023-0310221203321302"></a>

## Direct properties — sha256_digest / 302020022020 / 3

<a id="canonical-0221230202212112-0211131000331112-2321230213212000-3122310223312310-0210202320203302-2231231020101232-3230200313113132-3000021201202213"></a>

<a id="canonical-0313300332021330-2312301302332332-1200332231133102-3220002000022011-3020300213210301-1211110300231130-2311212103023111-0030213233003110"></a>

## digest property — sha256_digest / 302020022020 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-2231300121321322-2033231021331002-1012112030313203-1203030100233020-3022132313311122-1110211102000123-2021301100030010-2001202312121213"></a>

## Next pages — sha256_digest / 302020022020 / 5

- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-003.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2032220222001301-1021003012000102-0101002330300323-0312203302002020-0003200112112211-3021032200112002-2003312333133200-3113331200110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302221100020130-3330110313211132-0322013001310031-0101331211122021-3330032122300233-1332222121032313-3011003000203121-0110131222202302"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha384_digest — sha384_digest / 233311210021 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cds_record](resources--dns_zone--reference--group-002.md#canonical-0212131200323302-0300321013230033-2121301221232203-0233133001113103-3201313000002021-1131310322331320-1013231302130100-1033333322030133)
- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-003.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- primary.rr_set_group.rr_set.cds_record.values.sha384_digest

<a id="canonical-3120330220013230-0303113122202021-2201232202113323-2022113002210013-0330202233323122-0233122303011112-1210231220322032-0101310120112201"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha384 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
```

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

Terraform syntax:

```terraform
sha384_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131030031230003-1012001112231202-3213330232033231-1301310022303100-1211113001323011-0003022323033310-3201121331131332-3220202032322232"></a>

## Direct properties — sha384_digest / 233311210021 / 3

<a id="canonical-1130113223110200-0120022101323323-2031221211033022-0120123132011321-2021010302130322-0331110032010222-3300220222212311-0312312030112311"></a>

<a id="canonical-3311300032133131-0311113212200231-1023211203022102-0312031311320201-3311102301123231-1203120222312311-2221210320222222-3122331333003031"></a>

## digest property — sha384_digest / 233311210021 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(96, 96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-1312023313320202-3322213210011010-2221320331212003-2003223130221223-0301013220113313-0230311201031303-2322112010311312-1223001011310011"></a>

## Next pages — sha384_digest / 233311210021 / 5

- [primary.rr_set_group.rr_set.cds_record.values](resources--dns_zone--reference--group-003.md#canonical-3210332002003011-3212202303112030-0233233020201030-0213213201231121-1230302323101103-1121333022200310-2102321330133323-3033330230013033)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2122132002202221-2302210131231003-3030120300120123-2133331322231013-1130023111130312-3031311111321320-2202301111312003-0301302003020101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002112322021130-3101223213133010-3330003311021321-2331333110033012-2021211221201031-3322112002032032-1020203211223123-2222323203010023"></a>

## primary.rr_set_group.rr_set.cert_record — cert_record / 131200311102 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.cert_record

<a id="canonical-3012301320210302-1333331200213200-3010310033130210-1223002010211130-0022101013011131-0230313332121221-0332120101200133-1231213322232330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cert record.

Upstream description:

DNS CERT Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
cert_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010010213331333-1200030021302220-3321032003332222-1020002103203112-0030331122111103-3222313222222113-3032230313112223-2012012112333000"></a>

## Direct properties — cert_record / 131200311102 / 3

<a id="canonical-2020032320331031-0011123303232211-0321013130110322-0013222113013223-0111012212223311-1111032321332332-0222332031302331-0230333013301020"></a>

<a id="canonical-1113011122322320-0101222303200212-2332110012211002-3000132213300012-1033012010020120-3311230332021113-1112331020331133-1313130222113113"></a>

## name property — cert_record / 131200311102 / 4

Type: `"string"`. Optional.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-1110212302332301-1330131123332013-1122230122101023-3000000013013121-2213021011231101-1100133232033221-1123322331330012-3012232330233330): complete subsection reference.

<a id="canonical-0322103010211001-2023321212310303-1312033003000111-3112022303131110-1131131130113210-1202221123332023-0323100331132202-1010010133122321"></a>

## Next pages — cert_record / 131200311102 / 5

- [primary.rr_set_group.rr_set.cert_record.values](resources--dns_zone--reference--group-003.md#canonical-1110212302332301-1330131123332013-1122230122101023-3000000013013121-2213021011231101-1100133232033221-1123322331330012-3012232330233330)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1110212302332301-1330131123332013-1122230122101023-3000000013013121-2213021011231101-1100133232033221-1123322331330012-3012232330233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310220020233220-0333023320012021-3112111110113223-2302010130212201-0232331233122102-2212323122021200-1320001301222030-1312111132113010"></a>

## primary.rr_set_group.rr_set.cert_record.values — values / 333301310020 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--reference--group-003.md#canonical-2122132002202221-2302210131231003-3030120300120123-2133331322231013-1130023111130312-3031311111321320-2202301111312003-0301302003020101)
- primary.rr_set_group.rr_set.cert_record.values

<a id="canonical-2321003200110130-1233110210113031-1223210301301013-2121223122301131-1323123030313000-3023232220110202-0033103212123212-1003302122000220"></a>

Type: `"object"`. list nested block, Optional.

CERT Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("cert_key_tag",
    "certificate")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210322033311010-1220332333011331-2033011011302221-2321103221022121-0110100211002313-2121122020221312-1203111320220300-0133313233012032"></a>

## Direct properties — values / 333301310020 / 3

<a id="canonical-1211012122010310-0201120313120013-1032322211321110-3133121300311113-3210223231022230-1311321311232212-2301212110321012-1112220013310023"></a>

<a id="canonical-3223013032003100-0222301223120100-1021131300002001-1023023333302103-0321032002010021-0020013122303201-1313110323200230-2111033221011122"></a>

## algorithm property — values / 333301310020 / 4

Type: `"string"`. Optional.

\[Enum: RESERVEDALGORITHM|RSAMD5|DH|DSASHA1|ECC|RSASHA1ALGORITHM|INDIRECT|PRIVATEDNS|PRIVATEOID\]
CERT algorithm value must be compatible with the specified algorithm. - RESERVEDALGORITHM:
RESERVEDALGORITHM - RSAMD5: RSAMD5 - DH: DH - DSASHA1: DSASHA1 - ECC: ECC - RSASHA1ALGORITHM:
RSA-SHA1 - INDIRECT: INDIRECT - PRIVATEDNS: PRIVATEDNS - PRIVATEOID: PRIVATEOID. Possible values are
\`RESERVEDALGORITHM\`, \`RSAMD5\`, \`DH\`, \`DSASHA1\`, \`ECC\`, \`RSASHA1ALGORITHM\`, \`INDIRECT\`,
\`PRIVATEDNS\`, \`PRIVATEOID\`. Defaults to \`RESERVEDALGORITHM\`.

Upstream description:

CERT algorithm value must be compatible with the specified algorithm.

&#8203;- RESERVEDALGORITHM: RESERVEDALGORITHM

&#8203;- RSAMD5: RSAMD5

&#8203;- DH: DH

&#8203;- DSASHA1: DSASHA1

&#8203;- ECC: ECC

&#8203;- RSASHA1ALGORITHM: RSA-SHA1

&#8203;- INDIRECT: INDIRECT

&#8203;- PRIVATEDNS: PRIVATEDNS

&#8203;- PRIVATEOID: PRIVATEOID.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DH","DSASHA1","ECC","INDIRECT","PRIVATEDNS","PRIVATEOID","RESERVEDALGORITHM","RSAMD5","RSASHA1ALGORITHM"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "RESERVEDALGORITHM",
  "enum": [
    "RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1303010300120031-1323220122303030-2132211322123221-0022313212032000-0322201331020131-1120312020210320-2122311112312121-1212323022302322"></a>

<a id="canonical-3132201310332023-2230022211120212-3100300112102231-2111003031220222-3221333130212101-0312130022100001-3321001131222013-1323013200000332"></a>

## cert_key_tag property — values / 333301310020 / 5

Type: `"number"`. Optional.

Key Tag. Tag for categorization and filtering

Upstream description:

Tag for categorization and filtering

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2131303313320230-3121011203223201-3112232213321203-2021210323002222-3133020003302001-3033100333133332-3230300002002000-2210211112033030"></a>

<a id="canonical-2122332333131301-3233103202233220-2131323333201330-0010232013123221-2133010310010313-3200132010301203-3301102331212011-2301223220012202"></a>

## cert_type property — values / 333301310020 / 6

Type: `"string"`. Optional.

\[Enum: INVALIDCERTTYPE|PKIX|SPKI|PGP|IPKIX|ISPKI|IPGP|ACPKIX|IACPKIX|URI\_|OID\] CERT type value
must be compatible with the specified types. - INVALIDCERTTYPE: INVALIDCERTTYPE - PKIX: PKIX - SPKI:
SPKI - PGP: PGP - IPKIX: IPKIX - ISPKI: ISPKI - IPGP: IPGP - ACPKIX: ACPKIX - IACPKIX: IACPKIX -
URI\_: URI - OID: OID. Possible values are \`INVALIDCERTTYPE\`, \`PKIX\`, \`SPKI\`, \`PGP\`,
\`IPKIX\`, \`ISPKI\`, \`IPGP\`, \`ACPKIX\`, \`IACPKIX\`, \`URI\_\`, \`OID\`. Defaults to
\`INVALIDCERTTYPE\`.

Upstream description:

CERT type value must be compatible with the specified types.

&#8203;- INVALIDCERTTYPE: INVALIDCERTTYPE

&#8203;- PKIX: PKIX

&#8203;- SPKI: SPKI

&#8203;- PGP: PGP

&#8203;- IPKIX: IPKIX

&#8203;- ISPKI: ISPKI

&#8203;- IPGP: IPGP

&#8203;- ACPKIX: ACPKIX

&#8203;- IACPKIX: IACPKIX

&#8203;- URI\_: URI

&#8203;- OID: OID.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ACPKIX","IACPKIX","INVALIDCERTTYPE","IPGP","IPKIX","ISPKI","OID","PGP","PKIX","SPKI","URI_"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALIDCERTTYPE",
  "enum": [
    "INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1030030023212121-1310122023033021-1222333022110320-2030112330130132-2021320300013223-3231113123013013-0221132013100002-0330200333032321"></a>

<a id="canonical-1100113221301103-0023121020100023-0330110300230033-1212033023331102-3313133030313223-0122302120022202-3303300232310012-1130311000013300"></a>

## certificate property — values / 333301310020 / 7

Type: `"string"`. Optional.

Certificate. Certificate in base 64 format.

Upstream description:

Certificate in base 64 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3330130123121223-3232102303133202-3030032112001310-2321133330021233-0113330321232211-3212331230213100-0002011222212103-1213102301023101"></a>

## Next pages — values / 333301310020 / 8

- [primary.rr_set_group.rr_set.cert_record](resources--dns_zone--reference--group-003.md#canonical-2122132002202221-2302210131231003-3030120300120123-2133331322231013-1130023111130312-3031311111321320-2202301111312003-0301302003020101)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1103130121131032-2002012232031213-3122130221012010-2331212231230031-3110031301310331-0321330030021320-3212001113023221-1011201311323322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321230311311111-1110100300310031-2220021000022300-2211202113333322-2032220231002110-1330310100130213-3133031231000131-1213130012013201"></a>

## primary.rr_set_group.rr_set.cname_record — cname_record / 330200320230 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.cname_record

<a id="canonical-0310311121001122-2211200002312221-3310123210132123-1111003013102101-0220013300100011-3323021133223210-1203332212331011-2031212020103123"></a>

Type: `"object"`. single nested block, Optional.

DNSCNAMEResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
cname_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033120303031332-0321103023121021-0002030332330212-2321123112112310-3013132230223121-0101232033203023-0312231301003310-1000321121321332"></a>

## Direct properties — cname_record / 330200320230 / 3

<a id="canonical-1220111321122301-3001030230300300-0313021211213223-3313012203222220-3233002233010220-3300121311110312-2112210113022301-3022113321331133"></a>

<a id="canonical-3323030020030012-1033223133023220-1201301133330323-1101030331331333-2110231302101322-1221120333000211-1013030133202111-0131310313020130"></a>

## name property — cname_record / 330200320230 / 4

Type: `"string"`. Optional.

CName Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2120103131210132-3011322322030331-3011211302212322-0232133033012312-0020100212002222-2321013301033303-3231312232333112-3002112120212101"></a>

<a id="canonical-1123100313002203-3313220122311232-0220130130322110-2131012301022311-1011020030203320-2113303232022012-2020101200132311-1233322332212002"></a>

## value property — cname_record / 330200320230 / 5

Type: `"string"`. Optional.

Domain. Configuration parameter for value

Upstream description:

Configuration parameter for value

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-1222023303123113-2323211233220030-0000203032012100-0312122101203211-3303332102020320-3132200301320100-0232202310120033-1023203100010302"></a>

## Next pages — cname_record / 330200320230 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011222200100310-1010333112211321-3200202033332130-0300030030131011-2302032013132302-2131132331221122-2030312310210223-0013013203032003"></a>

## primary.rr_set_group.rr_set.ds_record — ds_record / 120320033200 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.ds_record

<a id="canonical-1312320202330202-3022222003331233-1020030132132200-1011023111332103-3330200202323012-1123022031311220-2221113301031332-2302121100021102"></a>

Type: `"object"`. single nested block, Optional.

DNS DS Record. DNS DS Record.

Upstream description:

DNS DS Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
ds_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122230103012333-3200131132222213-2233031030333312-1022213333030101-1122023202003111-3310313103211102-1130312212321023-1202303320302233"></a>

## Direct properties — ds_record / 120320033200 / 3

<a id="canonical-0032321310113023-1133320130332132-2121301000120130-1023313113323311-0330201003120100-2023003231303311-3101000322230333-1220332110021031"></a>

<a id="canonical-1113002033300230-0203300210303033-3210220203322222-0211200102023332-2120121213023302-0031033001023200-0010002112212012-2110221030311031"></a>

## name property — ds_record / 120320033200 / 4

Type: `"string"`. Optional.

DS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331): complete subsection reference.

<a id="canonical-0100023210222002-2100300333202333-1122312331300110-2332110212131223-0030322121123233-0322002200310310-0031022300300120-0313211233033200"></a>

## Next pages — ds_record / 120320033200 / 5

- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220210121330331-2223100011211011-3122012133011330-3313120011211203-3202302101030333-2332300223130132-2000020300122000-2022022331311320"></a>

## primary.rr_set_group.rr_set.ds_record.values — values / 210310221133 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- primary.rr_set_group.rr_set.ds_record.values

<a id="canonical-2100203210333023-2131000233103233-3110200332300003-0202103210212130-0311330213233211-3003301122200002-0210011133323020-3132332211132201"></a>

Type: `"object"`. list nested block, Optional.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key_tag"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha256_digest"),
  validators.ConflictingListObjectAttributes("sha1_digest",
    "sha384_digest"),
  validators.ConflictingListObjectAttributes("sha256_digest",
    "sha384_digest")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201312311201221-1313130300112002-1223132323212133-3220101020023113-0202333001102031-1011230211001213-1113322013023131-1321303003001030"></a>

## Direct properties — values / 210310221133 / 3

<a id="canonical-2323221310033131-0222331312210010-1322132223111200-0020031323310313-1203021201000031-0100113321133012-0303303100011131-1330230201310300"></a>

<a id="canonical-3331332320101000-2331101011200233-3310020123003302-3221301313200000-3310230313321020-3300132031111112-0100222332221231-2333213001001030"></a>

## ds_key_algorithm property — values / 210310221133 / 4

Type: `"string"`. Optional.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key-value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ECDSAP256SHA256","ECDSAP384SHA384","ED25519","ED448","RSASHA1","RSASHA1NSEC3SHA1","RSASHA256","RSASHA512","UNSPECIFIED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022013312210330-3111132222231013-0131033022321103-1203322330311311-0233320321200221-2103232330233210-1131331232030031-1303232033110231"></a>

<a id="canonical-0101131203211031-1131110313001113-0111201013103200-0231022330210322-0031022123331203-1313222123333000-2103233233000130-0111103011030203"></a>

## key_tag property — values / 210310221133 / 5

Type: `"number"`. Optional.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](resources--dns_zone--reference--group-003.md#canonical-1221213211021330-1020231010121110-1213003033310222-2002101033112223-2321201230033323-0103212031133023-0030002331221301-0111033300002012): complete subsection reference.

- [sha256_digest](resources--dns_zone--reference--group-003.md#canonical-1102203212111212-0030020311132130-0021222332121310-3011012022233211-0322312122333202-2332120213121001-0010201033302330-1100120312132210): complete subsection reference.

- [sha384_digest](resources--dns_zone--reference--group-003.md#canonical-1311123030010222-0013302212333220-2102221011302212-0333230023003222-3000320213321330-3113123310312010-3331130023103201-2223231323312320): complete subsection reference.

<a id="canonical-0002223000200012-3321213313103133-3213313010003332-0122122023233223-0111201331123010-0132311233303021-3103021213210312-1123313332331023"></a>

## Next pages — values / 210310221133 / 6

- [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](resources--dns_zone--reference--group-003.md#canonical-1221213211021330-1020231010121110-1213003033310222-2002101033112223-2321201230033323-0103212031133023-0030002331221301-0111033300002012)
- [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](resources--dns_zone--reference--group-003.md#canonical-1102203212111212-0030020311132130-0021222332121310-3011012022233211-0322312122333202-2332120213121001-0010201033302330-1100120312132210)
- [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](resources--dns_zone--reference--group-003.md#canonical-1311123030010222-0013302212333220-2102221011302212-0333230023003222-3000320213321330-3113123310312010-3331130023103201-2223231323312320)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1221213211021330-1020231010121110-1213003033310222-2002101033112223-2321201230033323-0103212031133023-0030002331221301-0111033300002012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220021012032001-3130123031020232-2132202002121122-1331123033031021-2031200220133013-2321212100323210-3001131300321013-1231320221122001"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha1_digest — sha1_digest / 212220000301 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- primary.rr_set_group.rr_set.ds_record.values.sha1_digest

<a id="canonical-0011133131321023-3130203301120032-0110112220131102-2111121011201021-2211331322120130-2002031220021330-2112231221322020-0223322123323200"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
```

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

Terraform syntax:

```terraform
sha1_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130210022203121-1022322303232121-1230202323221033-3301000011120121-0300113222321210-2230132320300121-0111311333012232-0132320120202203"></a>

## Direct properties — sha1_digest / 212220000301 / 3

<a id="canonical-2232131303232033-0333312022022312-1111110010021003-2101222120032030-3330123100230122-2130222210233312-0022103013111032-0202231013211033"></a>

<a id="canonical-0310202000233112-3222310121231002-0133012122331201-0323322031321101-2013330111013030-2300200322131230-2202330132331121-3003130101120023"></a>

## digest property — sha1_digest / 212220000301 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-0332112213313322-2112301121323032-0200022223001132-2301133203333012-2301033312102332-1332001123133110-3221230301220202-0222003310210333"></a>

## Next pages — sha1_digest / 212220000301 / 5

- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1102203212111212-0030020311132130-0021222332121310-3011012022233211-0322312122333202-2332120213121001-0010201033302330-1100120312132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030221320012002-1311213200202002-3331210201222023-0132032021220220-0032211112300031-2210132200312230-0130103013230200-3333133031232011"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha256_digest — sha256_digest / 322133200122 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- primary.rr_set_group.rr_set.ds_record.values.sha256_digest

<a id="canonical-3223300111103300-2131113003121202-0300202322223220-1212302211030213-1133121200002131-2300100110232331-0220122010230232-2213220212023120"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
```

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

Terraform syntax:

```terraform
sha256_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303112312202033-3033223320021123-1222233220201230-1230133123320331-2113121023122303-0103000000213130-3223202131212210-2000230012130033"></a>

## Direct properties — sha256_digest / 322133200122 / 3

<a id="canonical-1102232200130203-3213023331321211-0313023210120322-2121210303232303-3020110200013132-3312310210123331-3231002011333230-1213030122123000"></a>

<a id="canonical-3033202202030103-3103111032131321-3313231031030220-0233202100012122-1030303211300330-1210111020033231-1200012313131201-2330003112330031"></a>

## digest property — sha256_digest / 322133200122 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-0222113011100011-0032011020002012-0203012101201120-1320211310112002-1312120012121212-0103223122332332-3112322132012220-0102212033120323"></a>

## Next pages — sha256_digest / 322133200122 / 5

- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1311123030010222-0013302212333220-2102221011302212-0333230023003222-3000320213321330-3113123310312010-3331130023103201-2223231323312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133223123322230-3333112312031000-0320010333123200-2001200211120021-0132211110202311-3020131111000303-1122012232323003-2302112012103010"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha384_digest — sha384_digest / 100223103202 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.ds_record](resources--dns_zone--reference--group-003.md#canonical-0133111032113033-3001031330210022-3133010321131111-3030032331330223-1030103230301101-1222213021103231-1113333233030203-1300220320123223)
- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- primary.rr_set_group.rr_set.ds_record.values.sha384_digest

<a id="canonical-2020321033011002-0102000003310030-2202320000020333-0010133010030020-0130313010330210-3311003231003101-0212100220121010-2220132120202300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha384 digest.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("digest")}
```

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

Terraform syntax:

```terraform
sha384_digest {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101230003132312-0100302103112200-1033003000100213-3010222103012131-2213331000023132-1010131320232220-1000302002332012-3231302031220313"></a>

## Direct properties — sha384_digest / 100223103202 / 3

<a id="canonical-1320313210322131-2231002220303302-0332223110112201-2232123012321012-1221031322020202-0210103332300123-2223332212300301-1131002122200030"></a>

<a id="canonical-0303113302022312-1313330002101321-3322310201132223-2101113310112221-0220111033112132-1231001210122101-2023122020323221-1223122023233303"></a>

## digest property — sha384_digest / 100223103202 / 4

Type: `"string"`. Optional.

The 'digest' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(96, 96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-3333102033310133-1103223002002033-0332220121322131-2330301220003122-0230130030100201-2203312011300222-3001020031131212-0232322330113322"></a>

## Next pages — sha384_digest / 100223103202 / 5

- [primary.rr_set_group.rr_set.ds_record.values](resources--dns_zone--reference--group-003.md#canonical-3002100320102213-3300132232333231-3121011221223032-1102300223033120-0011311221003330-3021330103033101-0000320030230013-0310012032010331)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3030020123311102-0330002300121201-0201032021303132-0331310302301121-2002232330201131-3201010010133002-1121220200310131-0121101001003201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301221301212033-2310113210303003-1331033233012320-2131220020211113-0220230322032001-3010121320103031-0223133331320330-3113101230222003"></a>

## primary.rr_set_group.rr_set.eui48_record — eui48_record / 213231333212 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.eui48_record

<a id="canonical-0201123302222103-2312013113132012-0122000120310320-0320022121033100-3013022323220030-3331223213221333-2321013100233321-3302330110121232"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui48 record.

Upstream description:

DNS EUI48 Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
```

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

Terraform syntax:

```terraform
eui48_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310322123231103-1132330030023021-0221322022303130-0211030103331231-0100232310010333-1133223201210312-3232302130213022-0212103020030221"></a>

## Direct properties — eui48_record / 213231333212 / 3

<a id="canonical-2002122332210301-1330301002010310-3110211312012330-2323213122132332-1333031220133212-1220032013012131-0133100103102002-1331021133023020"></a>

<a id="canonical-3333202032322322-1132300210123010-3113221022011310-0231010110122001-2311322202021303-3230302221001033-1030203210123032-0031120210313022"></a>

## name property — eui48_record / 213231333212 / 4

Type: `"string"`. Optional.

EUI48 Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-0002023132231011-0303112102101232-1021233301002103-3320333213213013-1212111020313131-1311033201013113-2110200100221133-2121213220021303"></a>

<a id="canonical-3203301322300000-3120123221311021-1022321113312203-1120222311001012-3013201023110011-1132110213232003-3221020000113021-3020231310231131"></a>

## value property — eui48_record / 213231333212 / 5

Type: `"string"`. Optional.

EUI48 Identifier. A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Upstream description:

A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(17, 17),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 17,
  "minLength": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-3032012222033212-0302311300100110-0213301201202333-3233031013122211-3132213301123220-0320201201023003-3132213110211222-2130333203223012"></a>

## Next pages — eui48_record / 213231333212 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1123232112333312-0301023300332122-0100233233030001-1023002032221030-0013031220223001-2301310111033331-0120133033311203-0112320013203033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331211311221011-0012123223033310-1010001313202120-1132222132101020-3032102022120021-3331310010111232-0033010002033120-3223123123030101"></a>

## primary.rr_set_group.rr_set.eui64_record — eui64_record / 311003203010 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.eui64_record

<a id="canonical-2013211131032121-0312023111000010-0122020330121311-3122211222020033-0030010323222303-0111203231222112-2211112301312120-1210203033031233"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for eui64 record.

Upstream description:

DNS EUI64 Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("value")}
```

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

Terraform syntax:

```terraform
eui64_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322213320132210-2023122012112112-3331202122013330-0103011310330313-1201333230230020-1311232100303222-0003202331111133-3233220201233110"></a>

## Direct properties — eui64_record / 311003203010 / 3

<a id="canonical-3003013232120122-2111212112102312-2002331003000031-0223123230200221-1022202212233002-3101121031230310-3230022312231202-2221012103233032"></a>

<a id="canonical-2312200000112321-2111020031203232-1133113313200333-1033212223223310-3023133311303103-3333102221123003-0203111102221133-0223200003313310"></a>

## name property — eui64_record / 311003203010 / 4

Type: `"string"`. Optional.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-1102123323213102-3112220233132131-2232323223003302-0101010323201231-3130012311112111-3100222311132122-3113323010332001-1221321213220212"></a>

<a id="canonical-3121201313302011-2221110203203213-3133013001331230-3330102333112103-0001303222300133-3303032103203332-3120011300221101-3113312120110212"></a>

## value property — eui64_record / 311003203010 / 5

Type: `"string"`. Optional.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Upstream description:

A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(23, 23),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-1000031320200222-3210212211000131-2103000111032330-1302310323013013-3112100300213230-0232000201000020-2203101113120230-3121301000000212"></a>

## Next pages — eui64_record / 311003203010 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2011000133312033-0111302203233320-2010201212131212-2202122323010311-1322123301132002-1213120331011210-3311300321113201-3203231313323203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200122333301331-3101203222111302-3233331333010203-0320202013023300-0213203110301210-1320332033203110-1220200123212003-3310103303202300"></a>

## primary.rr_set_group.rr_set.lb_record — lb_record / 321320330223 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.lb_record

<a id="canonical-0021120310010231-0211130000112132-3000333101232313-3320012233310210-2300011022210313-2233120301332012-2023332030010130-0213202201302233"></a>

Type: `"object"`. single nested block, Optional.

DNS Load Balancer Record. DNS Load Balancer Record.

Upstream description:

DNS Load Balancer Record.

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

Terraform syntax:

```terraform
lb_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112112030000113-3122122022333112-3123011212113303-1012021223110332-2112330313123131-2131330010222100-0223133131032300-2221032020120120"></a>

## Direct properties — lb_record / 321320330223 / 3

<a id="canonical-0021020300011012-2022131312323132-3001202303113200-2131321121113030-1312131110310332-0120032231133222-1210312011100312-3321113221002200"></a>

<a id="canonical-1300323102310203-3130312030122222-3030110312022030-1231231313203101-3132201211133220-1323021230031010-0013321333012313-2123032301220332"></a>

## name property — lb_record / 321320330223 / 4

Type: `"string"`. Optional.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
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
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](resources--dns_zone--reference--group-003.md#canonical-2223311132322230-3320322312300132-0302230100002202-3110210102103332-0101321332303223-1231303031322303-3020002233313333-0103313232301202): complete subsection reference.

<a id="canonical-3002212032230110-2132302102003223-1100120230332201-0101013000333231-3200123120131200-0122313203330331-0210120030121222-0033121122213131"></a>

## Next pages — lb_record / 321320330223 / 5

- [primary.rr_set_group.rr_set.lb_record.value](resources--dns_zone--reference--group-003.md#canonical-2223311132322230-3320322312300132-0302230100002202-3110210102103332-0101321332303223-1231303031322303-3020002233313333-0103313232301202)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2223311132322230-3320322312300132-0302230100002202-3110210102103332-0101321332303223-1231303031322303-3020002233313333-0103313232301202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102223202010023-0221022110120123-1210012321122013-1202111321322300-1103022102332320-3223110233030302-0003012123132032-3132120000021320"></a>

## primary.rr_set_group.rr_set.lb_record.value — value / 023113230102 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--reference--group-003.md#canonical-2011000133312033-0111302203233320-2010201212131212-2202122323010311-1322123301132002-1213120331011210-3311300321113201-3203231313323203)
- primary.rr_set_group.rr_set.lb_record.value

<a id="canonical-2030002130231132-2000012332213113-1213331012213220-2121101123302313-2103300331310021-3120010212222111-2022230201301023-1002333113300020"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230322332032130-0122200030123102-0201202013120112-2321111033201120-0102213103013333-2030130211123311-0120122133331021-3021130310010010"></a>

## Direct properties — value / 023113230102 / 3

<a id="canonical-2130011331331021-0303222032023313-0222032030222131-1133331012200100-0032122011332321-3023221320120133-2130121033230223-0103230120303301"></a>

<a id="canonical-2010220232033002-1032202230032223-2310233130300200-1110032201323031-1311131211010233-3123102331033102-0302013302132023-3011302013233223"></a>

## name property — value / 023113230102 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2112210102230222-2032201010213103-3312310231113002-3130221302303220-2312300113013012-3023233212002331-0230213132012331-1230101230212321"></a>

<a id="canonical-2121300311301132-1203111332312130-1332001330111110-1021330103021100-3320011303103302-2132100313123131-0333203301030123-0231112121011203"></a>

## namespace property — value / 023113230102 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1322131212320031-0201330021321131-0313321131201233-1132132212021322-2000322002013013-1130310102311223-1033033120020102-1002010220201202"></a>

<a id="canonical-2302232123321103-0021113132220123-0331202312230002-0302101232111300-2012212310123330-3021132020211132-2003002323301002-3000100012013330"></a>

## tenant property — value / 023113230102 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0221133031011230-2000002202003201-0002033011311222-1223300221210233-2321000003303331-3303230100022333-1012100023101323-1311330201302300"></a>

## Next pages — value / 023113230102 / 7

- [primary.rr_set_group.rr_set.lb_record](resources--dns_zone--reference--group-003.md#canonical-2011000133312033-0111302203233320-2010201212131212-2202122323010311-1322123301132002-1213120331011210-3311300321113201-3203231313323203)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3003222202202221-1033200330333101-1301030003310233-1331212131201210-3302032301321211-0112023010221213-1200003133200302-0012102223210202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102021332003123-2313113003103232-0103332111211001-3010031113220220-0001002003130212-0010023011000022-2211300322321032-1010110110333100"></a>

## primary.rr_set_group.rr_set.loc_record — loc_record / 133000231233 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.loc_record

<a id="canonical-1203123000020330-2113200013101102-1113123013131013-1210020133113223-1223230311323133-1001123200011213-0103212102012311-0322023223102232"></a>

Type: `"object"`. single nested block, Optional.

DNS LOC Record. DNS LOC Record.

Upstream description:

DNS LOC Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
loc_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330330030011231-2020201030332100-0331222201113112-2021021030033222-3031130102202200-0220330132120112-1023001230000121-1113220000201131"></a>

## Direct properties — loc_record / 133000231233 / 3

<a id="canonical-2201011333233103-3121332211033331-2231300010202123-0332212313332201-2113003330013321-2100003331321020-0332101220101331-3102323322321131"></a>

<a id="canonical-2130111103010121-3213331320310232-0020000223313102-0322020032031022-2030110022313032-1233210102032210-3221103002011033-2212023312223230"></a>

## name property — loc_record / 133000231233 / 4

Type: `"string"`. Optional.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-1323013302112203-3032212023103013-1011322030210112-1133333031121233-3020121122120021-2303321121202030-3222123211002022-1012320101111312): complete subsection reference.

<a id="canonical-3321121200200101-0210012201101302-1130020011231303-2220121113332230-2330102020010211-3320222320331121-0221022111201211-3122332311331330"></a>

## Next pages — loc_record / 133000231233 / 5

- [primary.rr_set_group.rr_set.loc_record.values](resources--dns_zone--reference--group-003.md#canonical-1323013302112203-3032212023103013-1011322030210112-1133333031121233-3020121122120021-2303321121202030-3222123211002022-1012320101111312)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1323013302112203-3032212023103013-1011322030210112-1133333031121233-3020121122120021-2303321121202030-3222123211002022-1012320101111312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012110001022013-1010001002110313-2111103032323332-0323331223030330-2112033221233312-2211100110203230-2330030110230330-3323322033220230"></a>

## primary.rr_set_group.rr_set.loc_record.values — values / 113210201113 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--reference--group-003.md#canonical-3003222202202221-1033200330333101-1301030003310233-1331212131201210-3302032301321211-0112023010221213-1200003133200302-0012102223210202)
- primary.rr_set_group.rr_set.loc_record.values

<a id="canonical-3313311200021130-1131002203100230-2031300301002002-3101101210022010-3311213223203013-2322301230200210-3020231212101013-2213131332202110"></a>

Type: `"object"`. list nested block, Optional.

LOC Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("altitude",
    "latitude_degree",
    "longitude_degree")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320131330110320-3211001112212232-2333312131113110-1131011011103030-1212100032313122-0122322320033201-1131331130033113-3013001321102301"></a>

## Direct properties — values / 113210201113 / 3

<a id="canonical-0013112313120332-3320132333222132-1331001033001122-0033001311322013-0302211032210011-3131023033300010-3000323211033002-0313030203101130"></a>

<a id="canonical-3322313010213013-2323111131222303-0033220011221323-3201303210311031-2121122133002100-1312201231133122-0003223310312221-3112122303220013"></a>

## altitude property — values / 113210201113 / 4

Type: `"number"`. Optional.

Altitude. Altitude in meters.

Upstream description:

Altitude in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0302000232331101-3133201211201122-3132233113031112-2323313130202220-3321020231222211-2130033321232001-0012111322120301-0313020120221001"></a>

<a id="canonical-2120301333321012-2131030301031233-0030223310030023-0331303201121200-2321220133112100-0211110200100021-2032110232201101-3330230000200011"></a>

## horizontal_precision property — values / 113210201113 / 5

Type: `"number"`. Optional.

Horizontal Precision. Horizontal Precision in meters.

Upstream description:

Horizontal Precision in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-3022310322123313-0311101020222012-1020230200300031-2322110031022310-2102101131200133-0132113323121131-3310021130310013-1123200031113211"></a>

<a id="canonical-3013322230131113-3300110311021302-2223231223000023-0231030233131301-3232011320121122-2311222021202122-3303103202111013-0100010312312032"></a>

## latitude_degree property — values / 113210201113 / 6

Type: `"number"`. Optional.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 90),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3310130223121023-2221330010131111-2001312003111220-0232232312132011-1310113100303011-1321212133120332-2030301033020023-3123133000211103"></a>

<a id="canonical-2012222002210100-2310013131021301-3130222032103132-2010223111112111-1101210303320301-3111211302130202-0210100011230213-1323201223310321"></a>

## latitude_hemisphere property — values / 113210201113 / 7

Type: `"string"`. Optional.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Upstream description:

Latitude hemisphere can only be N or S

&#8203;- N: North Hemisphere

&#8203;- S: South Hemisphere.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["N","S"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("N",
    "S"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000212200110013-0000201320333332-0113331020020110-2102202223021330-2132110102021133-1310031021020032-3133313213032313-2330021201020213"></a>

<a id="canonical-3111010121300333-3023301303030332-0013020301232133-0302231022313132-3020221331330010-2222032120011330-2121001300120023-2232100330313200"></a>

## latitude_minute property — values / 113210201113 / 8

Type: `"number"`. Optional.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-0321123302220223-1213203131102300-1000112213311010-3302322030020202-0111130331232310-0002103001322223-1112231210023113-1223123312232322"></a>

<a id="canonical-3302323021312013-0123100021032230-2300321112130323-0101023000212032-3201130221213303-1003213100111311-2101031231210311-3332000302302310"></a>

## latitude_second property — values / 113210201113 / 9

Type: `"number"`. Optional.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-0000130111101322-2233132231333001-2302120220310101-3312021312221011-0112012322212233-2003103030011131-0030222122223300-2132111303102301"></a>

<a id="canonical-3111113020223033-0112303333322330-0130002001102020-3122000223010111-1201111232021302-2333101120111023-3211331130010310-2102201301202033"></a>

## location_diameter property — values / 113210201113 / 10

Type: `"number"`. Optional.

Diameter of a sphere enclosing the described entity, in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-1112023302312232-1331003021013200-1021302200332101-2110023122323012-2123311211120002-0332122102330131-3121320123011311-3020232223323122"></a>

<a id="canonical-0000013322213310-1113300033310201-2110133312022000-0122310011011202-0232121331123132-1121313323131200-2323133230230001-3231133231020011"></a>

## longitude_degree property — values / 113210201113 / 11

Type: `"number"`. Optional.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 180),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3311213031211000-2203110013110303-3011131310313213-2102122010222302-0311321031303203-3303003111302301-1221312122122210-3213102101323012"></a>

<a id="canonical-2033210000213122-1301120210312330-1101230022320020-0213032003021103-2111000102200032-2031030110123233-3110011230210330-2002221032212322"></a>

## longitude_hemisphere property — values / 113210201113 / 12

Type: `"string"`. Optional.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Upstream description:

Longitude hemisphere can only be E or W

&#8203;- E: East Hemisphere

&#8203;- W: West Hemisphere.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["E","W"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("E",
    "W"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3333311133103112-1223122203322303-1101213130001103-2003122121110130-0121201213333112-0302012011133303-3302000002003231-0010030132101310"></a>

<a id="canonical-2231131000002132-2031221313120212-0303131213012230-2110322110311023-2313030303320132-0032123023131012-3202233133022202-1312313222232202"></a>

## longitude_minute property — values / 113210201113 / 13

Type: `"number"`. Optional.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 59),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-0003300311211321-3302033301013000-1303223131102001-0021310012301222-3002221110003121-3102100110322301-3301103010223123-0202202000321131"></a>

<a id="canonical-0131322120110032-1011020313312122-1333133323320222-2303102310232100-1333013001030213-0003120022320213-1022212103233312-1000220120111102"></a>

## longitude_second property — values / 113210201113 / 14

Type: `"number"`. Optional.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-3023332133221031-2213200201011103-3213012023033333-0312213013012003-2200302200332113-0320122133201000-3000300330100222-0030211201320001"></a>

<a id="canonical-0110002211230030-2012310313011110-2213122012103203-0003232311200200-3001213013311213-1113122213110200-3110202100210022-3333322300111102"></a>

## vertical_precision property — values / 113210201113 / 15

Type: `"number"`. Optional.

Vertical Precision. Vertical Precision in meters.

Upstream description:

Vertical Precision in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-1030203323101303-3113303323231031-2011121023120231-3310202022223200-0221120002010321-1111120300312122-1013110002300323-1303122200232310"></a>

## Next pages — values / 113210201113 / 16

- [primary.rr_set_group.rr_set.loc_record](resources--dns_zone--reference--group-003.md#canonical-3003222202202221-1033200330333101-1301030003310233-1331212131201210-3302032301321211-0112023010221213-1200003133200302-0012102223210202)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2003001322011302-1301231333320300-1203302312221330-2033303022131132-3121320130110031-2332100310010333-2130011113103032-1022320310102202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122020202222122-3200033213010313-3211200032310000-0230313112203201-0032120202032121-1201232010310033-0302003001012100-3311112132232133"></a>

## primary.rr_set_group.rr_set.mx_record — mx_record / 210303032203 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.mx_record

<a id="canonical-2132030313003332-3230003011110221-0130002031232112-3133020121102322-2223230102301312-2023030330230330-3131300100100313-3221132223101331"></a>

Type: `"object"`. single nested block, Optional.

DNSMXResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
mx_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223001123303232-2311302102232033-2032312212301312-1021231003231312-3323313030131300-1102022303221323-1012112023021212-0222232302130230"></a>

## Direct properties — mx_record / 210303032203 / 3

<a id="canonical-0102031200201131-1132303302003222-3210331133220110-3332322020300030-0033310330203230-2010332323122023-0032030122310211-1112003010201310"></a>

<a id="canonical-0331002223330212-2022210210120230-2101302220033132-2113122321230001-1221233323212112-2220232322122323-1023112330013332-1100011211002022"></a>

## name property — mx_record / 210303032203 / 4

Type: `"string"`. Optional.

MX Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-3203011211313001-0002133210230122-3102011130012031-2123133001220331-1333031133312300-1102233313120122-3020331320013120-3302110102201210): complete subsection reference.

<a id="canonical-2131131132221232-2303111130310102-0110023101330011-1112110003232310-2333300233232322-3033011113012333-3200331222000120-2133112131201321"></a>

## Next pages — mx_record / 210303032203 / 5

- [primary.rr_set_group.rr_set.mx_record.values](resources--dns_zone--reference--group-003.md#canonical-3203011211313001-0002133210230122-3102011130012031-2123133001220331-1333031133312300-1102233313120122-3020331320013120-3302110102201210)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3203011211313001-0002133210230122-3102011130012031-2123133001220331-1333031133312300-1102233313120122-3020331320013120-3302110102201210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112200001220210-2121121231303102-0021110000130213-2323101201222231-1023130330321001-1133331023220311-2013112302031000-0211303322101231"></a>

## primary.rr_set_group.rr_set.mx_record.values — values / 122022123022 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--reference--group-003.md#canonical-2003001322011302-1301231333320300-1203302312221330-2033303022131132-3121320130110031-2332100310010333-2130011113103032-1022320310102202)
- primary.rr_set_group.rr_set.mx_record.values

<a id="canonical-2202333202101131-1132102103013000-1112330121332320-3212223020221231-2023130231102331-3203303102210210-0033220002221200-0313132030222011"></a>

Type: `"object"`. list nested block, Optional.

MX Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001120210103100-0122300202300200-3003302122022122-2100032032112130-1233101123231200-1130302013321301-0323113131012023-3003320232333220"></a>

## Direct properties — values / 122022123022 / 3

<a id="canonical-1133102322110311-3101323123323131-3101301203111231-3223212323211213-1331110102101010-3233312022332330-2030332033303201-1102021310032131"></a>

<a id="canonical-2112012012212323-1011223221010201-1330032313123333-0133123302011320-2210023230213300-0312330312202033-2011231311220200-2321022321023020"></a>

## domain property — values / 122022123022 / 4

Type: `"string"`. Optional.

Mail exchanger domain name, please provide the full hostname, for.

Upstream description:

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3111313022123021-1133213322112032-2203121002200232-1121301000133330-3000133312233103-1333132030312300-1102030223123323-0022112233333301"></a>

<a id="canonical-3111131313001002-2313130002110011-1001032131130133-3120232111312110-0131202110302232-0023032223031222-2002013232002132-0130003111001031"></a>

## priority property — values / 122022123022 / 5

Type: `"number"`. Optional.

Priority. Mail exchanger priority code.

Upstream description:

Mail exchanger priority code.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0222301231021021-2321310313012320-3210230130112031-2302221130033300-0323022021011200-1300121231311322-1133033110021321-1101030230333033"></a>

## Next pages — values / 122022123022 / 6

- [primary.rr_set_group.rr_set.mx_record](resources--dns_zone--reference--group-003.md#canonical-2003001322011302-1301231333320300-1203302312221330-2033303022131132-3121320130110031-2332100310010333-2130011113103032-1022320310102202)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0031001220121130-2232011132130122-3223223310210021-2102322233011223-2233030012200001-3232113133123013-2230020011101011-3203003021223223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121111313002311-2101111110301202-2232110213020131-2130000321003001-1201222310210113-3203121023112103-1323012223321200-3322220212332022"></a>

## primary.rr_set_group.rr_set.naptr_record — naptr_record / 322023012120 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.naptr_record

<a id="canonical-0113102211212002-1101222021022110-3101000310021221-3213123131122132-2222303310010210-3123021310123133-0312033331022113-3003012030033011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for naptr record.

Upstream description:

DNS NAPTR Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
naptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122120103101213-3320102131302331-1332313122301220-0132201320012221-0321112113000023-1133320201312320-2133331201201330-3311320233011202"></a>

## Direct properties — naptr_record / 322023012120 / 3

<a id="canonical-1200012232231001-1301100312311031-2133300113100220-1030101021010222-2032223233100320-0333201103130110-0123103332311012-3120002221030012"></a>

<a id="canonical-3301220122323002-0113001330213122-1221311131023011-2231112333332100-1200303131213301-0020212003230222-0111302300323032-2200200313102013"></a>

## name property — naptr_record / 322023012120 / 4

Type: `"string"`. Optional.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-0100220222133303-1230232023133312-3132333010213303-0000021310121331-0121023031303003-2003023202002101-3320001002112000-1133232110312031): complete subsection reference.

<a id="canonical-1010311020002313-0132310202200131-0131000332103121-2120123312233002-1333222221013333-0200020012121123-0122000131032100-3310100100130301"></a>

## Next pages — naptr_record / 322023012120 / 5

- [primary.rr_set_group.rr_set.naptr_record.values](resources--dns_zone--reference--group-003.md#canonical-0100220222133303-1230232023133312-3132333010213303-0000021310121331-0121023031303003-2003023202002101-3320001002112000-1133232110312031)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0100220222133303-1230232023133312-3132333010213303-0000021310121331-0121023031303003-2003023202002101-3320001002112000-1133232110312031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022320311320031-1300120011010013-3232321003033010-2132331011112131-2303030320101110-0303303131033320-0013330301033113-3313010120123021"></a>

## primary.rr_set_group.rr_set.naptr_record.values — values / 232023012322 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--reference--group-003.md#canonical-0031001220121130-2232011132130122-3223223310210021-2102322233011223-2233030012200001-3232113133123013-2230020011101011-3203003021223223)
- primary.rr_set_group.rr_set.naptr_record.values

<a id="canonical-0222200233003323-0322031231313100-1201022110113322-1130023112121030-3103233101313121-1313030302302333-2322013132210231-0032111113103033"></a>

Type: `"object"`. list nested block, Optional.

NAPTR Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("flags",
    "order",
    "preference")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201333103133120-1302131002332330-2200312223320331-2211223221322010-3131323201013213-3320313221120201-2313212112223002-2211212311120221"></a>

## Direct properties — values / 232023012322 / 3

<a id="canonical-2000332132310023-1130303122313003-2101233300213131-2030002031201201-3111212121131033-1302032203222301-3013111030321123-0112300300330223"></a>

<a id="canonical-3131323232010220-0300033020330113-0211123133203211-0332121110221120-3321133202100231-0311133132221013-0023020311210333-0001332123210123"></a>

## flags property — values / 232023012322 / 4

Type: `"string"`. Optional.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="canonical-1131013031102231-1311302210303031-1031033032330023-0013330201133112-1300021322013100-2210010012012122-3132301203200123-0231312001211031"></a>

<a id="canonical-3011300303002113-3010311223032000-2022010013110133-2103332022113102-0102200220313212-0120021200003233-2223221331202111-0202302131320323"></a>

## order property — values / 232023012322 / 5

Type: `"number"`. Optional.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2011121233210300-0121201122303221-3300003023003222-0132232213132022-2310301231212110-1203320320310023-2113033110111312-1203321032011212"></a>

<a id="canonical-1231330021122112-3011101000110202-2222210001103003-1023330232322123-0132012303233213-1302310011100032-2313313220111121-1211303032122031"></a>

## preference property — values / 232023012322 / 6

Type: `"number"`. Optional.

Preference when records have the same order. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1022122030201233-0030103013313101-1220222121123303-1230222122133123-2112013030212201-3332121011100211-0330212323001202-3302211120230300"></a>

<a id="canonical-3321232322231220-3123032230323321-0311020312302312-0132321333001102-0210332313023133-3121302110333332-3221020121032133-2233300120001123"></a>

## regular expression property — values / 232023012322 / 7

Type: `"string"`. Optional.

Regular expression to construct the next domain name to lookup.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-1202323213312222-3123130121131000-3331022313122333-2223222230203203-1020000320101032-3102220101303203-1212203303132201-1211223210113211"></a>

<a id="canonical-1322310302223120-3132211103112003-3032232303021302-3133311023031321-0032132220132311-0312023331023022-2203031100220003-1023023020230331"></a>

## replacement property — values / 232023012322 / 8

Type: `"string"`. Optional.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1030011011201021-2321133132211031-2303023321320202-1201311121123320-2320202102312212-1333111131112123-2131131301202323-1123331311213000"></a>

<a id="canonical-2200201301232131-1211010022221302-1022312202200032-1130131120032302-3101313300013110-3321200101031110-0112100131023202-2222312120013231"></a>

## service property — values / 232023012322 / 9

Type: `"string"`. Optional.

Specifies the service(s) available down this rewrite path.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```

<a id="canonical-3121120231300032-0123323213112311-0231212310313002-1002003010102212-3103320303022301-2032202323232333-2213112202020113-0333303312033030"></a>

## Next pages — values / 232023012322 / 10

- [primary.rr_set_group.rr_set.naptr_record](resources--dns_zone--reference--group-003.md#canonical-0031001220121130-2232011132130122-3223223310210021-2102322233011223-2233030012200001-3232113133123013-2230020011101011-3203003021223223)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3112102233211200-0313123203112300-1130013102230113-2011023232120220-1302012230210220-3202131201133110-0011021300201313-1102311332102132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100301123130020-3302220311011323-2310302223223013-3300121033133313-2223013333021311-2222110122020131-0233211200222033-1011310010003031"></a>

## primary.rr_set_group.rr_set.ns_record — ns_record / 020202032233 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.ns_record

<a id="canonical-1231003122123303-2133013001222111-3132233130013001-3012311102221220-2103210112001212-3323000203323113-3120302331000232-2021211213312321"></a>

Type: `"object"`. single nested block, Optional.

DNSNSResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
ns_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212130032303313-2221001321333220-2320203313132110-3000033022120133-3212220131001221-2303122122321301-2122103303210212-1202003031023312"></a>

## Direct properties — ns_record / 020202032233 / 3

<a id="canonical-3320020101210010-1122130010031303-3301323332322213-2301330010001022-0121210231130222-1032122301130321-1320031201003011-2112231102302113"></a>

<a id="canonical-3100002110303233-1000131003212001-0323323322003303-3030232313122002-0112210322110301-3332120000120321-2302001332033220-3211302110332021"></a>

## name property — ns_record / 020202032233 / 4

Type: `"string"`. Optional.

NS Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2013013321002000-1221302210011312-1303231000120231-3302010111320010-1331010332113222-2302333313221033-3033000111223302-1212201201230302"></a>

<a id="canonical-3022000103022023-0213310203101132-1013111233002211-3230331011003022-2312031022330010-1232031101011332-1102333111230120-2120123320233323"></a>

## values property — ns_record / 020202032233 / 5

Type: `["list", "string"]`. Optional.

Name Servers. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0332000322302033-3310112111102210-3022333323202233-0231123212202302-3301221133120011-1223303202122312-3203303212212101-0203210303013303"></a>

## Next pages — ns_record / 020202032233 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3232000310031110-1111032130320130-1330010220023223-2023200213300002-3123322311111130-2121011102311023-1203013003000322-3300000321210110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200232120201130-3023122202210111-0021131213323022-2310202220311210-2121113111021231-2012123010301031-2312012002303332-3111100220303323"></a>

## primary.rr_set_group.rr_set.ptr_record — ptr_record / 000002332013 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.ptr_record

<a id="canonical-3013221023311022-0020211130201303-3131202122033003-1300210133223000-3222010023132110-1001102031230113-0020222130133000-1122303200203200"></a>

Type: `"object"`. single nested block, Optional.

DNSPTRResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
ptr_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003003020300201-1203300332302302-2211313302121102-2201013003222122-0303220032222212-1332012001133103-0012002320012212-0220220120300033"></a>

## Direct properties — ptr_record / 000002332013 / 3

<a id="canonical-3100223311113131-0103210012013210-0331322110221133-0020032321022320-2022203200101222-2310003322230001-0213311230111011-1202012120212032"></a>

<a id="canonical-0133233032131101-1331210130310230-0321321100231213-2230012003332300-1232321122032222-3300301232110012-2113230311200201-0003202223123301"></a>

## name property — ptr_record / 000002332013 / 4

Type: `"string"`. Optional.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2023312301232211-1132111221331013-2102012300332121-0312000310311303-0030132032012211-3231013321332002-1203032112210230-0230333012212213"></a>

<a id="canonical-0000011133310331-2331223123110032-3110322210000233-3222020002213311-2001313313133320-2032321110210302-3031323131123330-1231312021021100"></a>

## values property — ptr_record / 000002332013 / 5

Type: `["list", "string"]`. Optional.

Domain Name. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1330330133010001-2111031031231323-0123321103000001-2101121300310230-2031130321312022-2322021112231332-3301000020133210-1011011211110332"></a>

## Next pages — ptr_record / 000002332013 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3331101323210103-3033131313122302-0130321211322012-2233211131121323-1203222333302212-2332301313331233-3123010110333331-2023103210211121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010011213222131-1303312101320110-2122021302102011-1010210011031233-1121132211230010-3211103102001123-2022032031312023-0331210333231300"></a>

## primary.rr_set_group.rr_set.srv_record — srv_record / 323002000120 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.srv_record

<a id="canonical-2001132300321330-0200133330021203-0313213322013100-0111121110201021-2012001100220230-1000022333221202-3112310302033103-1021221233100113"></a>

Type: `"object"`. single nested block, Optional.

DNSSRVResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "values")}
```

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

Terraform syntax:

```terraform
srv_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010210300230123-1121330222223200-3202201032220311-3223102123021231-2302112212321312-0201321002303203-0231131011220301-0113023100222122"></a>

## Direct properties — srv_record / 323002000120 / 3

<a id="canonical-3022131012200113-3123121211100100-0221003010123321-0213303131121032-3122323030123030-0320131123002113-3101323220311231-2330310322133201"></a>

<a id="canonical-1033122303101102-0223130022033213-0322321202210220-2000112010232332-3200222030112320-3212110203123330-2333230113120231-3321211021332201"></a>

## name property — srv_record / 323002000120 / 4

Type: `"string"`. Optional.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-1313232132213222-3323312131222220-1302231230301310-2210311220110031-0003120131301303-3032221203121320-2213022201330021-3223010010123021): complete subsection reference.

<a id="canonical-0302313202121210-1131122131200021-0100023121303031-2112330322120000-1111230011033003-3200100222111230-1032203021133211-3310213332231131"></a>

## Next pages — srv_record / 323002000120 / 5

- [primary.rr_set_group.rr_set.srv_record.values](resources--dns_zone--reference--group-003.md#canonical-1313232132213222-3323312131222220-1302231230301310-2210311220110031-0003120131301303-3032221203121320-2213022201330021-3223010010123021)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1313232132213222-3323312131222220-1302231230301310-2210311220110031-0003120131301303-3032221203121320-2213022201330021-3223010010123021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010332321210210-3301110201001323-1231303223131022-3133233023211320-3223130113323100-1121122020300131-3220332010301120-1203132231231011"></a>

## primary.rr_set_group.rr_set.srv_record.values — values / 001223003013 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--reference--group-003.md#canonical-3331101323210103-3033131313122302-0130321211322012-2233211131121323-1203222333302212-2332301313331233-3123010110333331-2023103210211121)
- primary.rr_set_group.rr_set.srv_record.values

<a id="canonical-0330111210022221-0223330313120313-0123322001022203-3102121300120123-2000122203102101-1130011312000333-2023210331301111-1111113130321111"></a>

Type: `"object"`. list nested block, Optional.

SRV Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100200121000233-2210102331010232-2101023020230123-0212322302333331-3212000110010230-3322303331332212-0000211012132133-2320211123222001"></a>

## Direct properties — values / 001223003013 / 3

<a id="canonical-0320200121330033-2102220303010332-2331133001210202-0122123203111103-3221312301130130-1333222312220221-1110032222223301-3010223032320201"></a>

<a id="canonical-0200212333023121-1212111203231223-2312323233230222-2200012023010321-0330320001121122-0220122121230100-0131021223200221-3223301103211130"></a>

## port property — values / 001223003013 / 4

Type: `"number"`. Optional.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0330033232222012-2031220323210103-0232222212023222-1202232323320001-0113231211002301-3321331300333002-1132313110213332-2303222323230010"></a>

<a id="canonical-1300101200310120-2221121033322233-1220100033003103-0132300010001023-0100032010100112-0310020031121313-0013321212301030-0002022020313103"></a>

## priority property — values / 001223003013 / 5

Type: `"number"`. Optional.

Priority of the target. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0230233002211220-0231113232232120-0113320001010100-3301022031010221-2213213123333233-3202101320301203-0310330333313131-2122120102222023"></a>

<a id="canonical-0011133300023321-2232212213121011-3022003002312132-2331200120332233-0323123023213032-3020220223230223-1322002102201302-2203323101013321"></a>

## target property — values / 001223003013 / 6

Type: `"string"`. Optional.

Hostname of the machine providing the service.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-1111203100210131-0200111203203233-3202120003030231-0203121301223000-0202113011300021-3211321033111023-3301200001303300-1312311121000120"></a>

<a id="canonical-3303303022033110-3100023022201020-0221200332303121-3302202030012120-2101230103121300-0131002020333123-3213232100300330-2313302102300102"></a>

## weight property — values / 001223003013 / 7

Type: `"number"`. Optional.

Weight of the target. A higher number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0130032022000002-0322203302223210-1032232213123111-1232302013233010-3013112321322233-3323010330320123-0202100223111230-2300213332323101"></a>

## Next pages — values / 001223003013 / 8

- [primary.rr_set_group.rr_set.srv_record](resources--dns_zone--reference--group-003.md#canonical-3331101323210103-3033131313122302-0130321211322012-2233211131121323-1203222333302212-2332301313331233-3123010110333331-2023103210211121)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233221330020202-2111121302011233-1310113101030013-1113103313013230-2101110331201312-1321011023312103-3332212230031023-0033211033203333"></a>

## primary.rr_set_group.rr_set.sshfp_record — sshfp_record / 013111001031 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.sshfp_record

<a id="canonical-0121121322011211-1303301211230003-0201033312000112-0101223100102212-2102100201221222-1233012320013113-2021333323302131-0213331312123121"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sshfp record.

Upstream description:

DNS SSHFP Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
sshfp_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122011000212322-1013231220010200-1222013011211022-0220001130131003-2132020203020101-0202000012031112-0000323000213311-3133203323222230"></a>

## Direct properties — sshfp_record / 013111001031 / 3

<a id="canonical-1213032112302330-2022300300122131-3030013113030212-0301021202311320-1331110100322013-3112223312103322-1201011233013001-1013022130133033"></a>

<a id="canonical-3012013203020321-3013220112212013-1020233232320201-0001003321202111-0112010012331233-2203223121223022-3011033020133202-3323013030112020"></a>

## name property — sshfp_record / 013111001031 / 4

Type: `"string"`. Optional.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322): complete subsection reference.

<a id="canonical-3021301210210020-3030113012113021-1322330301220221-2013222313031202-1203330221120010-0331212100000013-1100000031211300-0121020200311213"></a>

## Next pages — sshfp_record / 013111001031 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102131103301222-1331101311233230-1320321203131320-0221313310232013-0333322203010231-3332133130231003-2302223211022003-2021321020220202"></a>

## primary.rr_set_group.rr_set.sshfp_record.values — values / 231120002011 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311)
- primary.rr_set_group.rr_set.sshfp_record.values

<a id="canonical-3300100220311322-3102300013023303-1002222232103113-1010021032220123-3021101232031310-0113110031110203-1030020300032022-2311333323323231"></a>

Type: `"object"`. list nested block, Optional.

SSHFP Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("sha1_fingerprint",
    "sha256_fingerprint")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313310233303013-2032023121032112-3120022330220123-3102203220301203-0231022211232011-3330231030310110-1013101010312233-2022112300022022"></a>

## Direct properties — values / 231120002011 / 3

<a id="canonical-2002221222101033-0010212011212120-3123322310313233-3331213133311320-1023101333213032-3320100113110003-2022031113231031-1121332202330311"></a>

<a id="canonical-1112110323211301-2210232033200320-3012333002031331-1010313121111233-0203000121321213-0211302122210100-1111210130300320-2223300012231102"></a>

## algorithm property — values / 231120002011 / 4

Type: `"string"`. Optional.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Upstream description:

SSHFP algorithm value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM

&#8203;- RSA: RSA

&#8203;- DSA: DSA

&#8203;- ECDSA: ECDSA

&#8203;- Ed25519: Ed25519

&#8203;- Ed448: Ed448.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["DSA","ECDSA","Ed25519","Ed448","RSA","UNSPECIFIEDALGORITHM"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIEDALGORITHM",
  "enum": [
    "UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [sha1_fingerprint](resources--dns_zone--reference--group-003.md#canonical-2323200133101212-1233020310330331-1020330120003130-0301212011321313-0033033220010232-3103312203323111-2203311212112222-2000323133201001): complete subsection reference.

- [sha256_fingerprint](resources--dns_zone--reference--group-003.md#canonical-3103002202322302-2232010213112211-1102001223032112-3101301022320121-0102223210112131-3030021311311121-1133202232131312-2302110021303323): complete subsection reference.

<a id="canonical-0121230300323113-3003111203300201-1320012102030220-2021120232111312-1222012321002121-1030131221312222-2322021202331100-1101213031303303"></a>

## Next pages — values / 231120002011 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](resources--dns_zone--reference--group-003.md#canonical-2323200133101212-1233020310330331-1020330120003130-0301212011321313-0033033220010232-3103312203323111-2203311212112222-2000323133201001)
- [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](resources--dns_zone--reference--group-003.md#canonical-3103002202322302-2232010213112211-1102001223032112-3101301022320121-0102223210112131-3030021311311121-1133202232131312-2302110021303323)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2323200133101212-1233020310330331-1020330120003130-0301212011321313-0033033220010232-3103312203323111-2203311212112222-2000323133201001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311120202111130-0331131332133223-2330301013220131-0013231231221123-1333231120313112-0211001323313311-1013112101311021-1110322031003231"></a>

## primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint — sha1_fingerprint / 131123130211 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311)
- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322)
- primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint

<a id="canonical-1130201023023123-3003233001022022-1103002023221203-3120103301012110-2002112000301333-0001111201012033-0332131013022210-3321213103103321"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha1 fingerprint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
```

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

Terraform syntax:

```terraform
sha1_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122210021333022-1212230121312321-2322132210030320-0030313222332113-1200203111033100-0200210232302120-0322132133020101-0003020301112121"></a>

## Direct properties — sha1_fingerprint / 131123130211 / 3

<a id="canonical-3012211102132302-0032303133312012-3112122000300033-3333100023133023-0333322332032103-0010100000331323-1101213233331000-0322311032202022"></a>

<a id="canonical-2210202222330120-1232013013312301-0212231202301130-3200232321201303-3201300320121001-0311200111321111-3330323312003200-3230101103013023"></a>

## fingerprint property — sha1_fingerprint / 131123130211 / 4

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(40, 40),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-2300123210001211-3230000110021112-2202002230233320-0031312210111303-0301302220311003-2201123203230311-3122311321100100-1212132203121121"></a>

## Next pages — sha1_fingerprint / 131123130211 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3103002202322302-2232010213112211-1102001223032112-3101301022320121-0102223210112131-3030021311311121-1133202232131312-2302110021303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200203131103221-1003123122312030-0010320203133130-2212002300023001-0003313302112123-0010103101330233-2212102321220321-0313330330131013"></a>

## primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint — sha256_fingerprint / 013200211003 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.sshfp_record](resources--dns_zone--reference--group-003.md#canonical-3021303131230121-2130011003320213-1320312132022020-1030220220200211-1220030223000022-0033023310312020-1222101103301213-2133330022031311)
- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322)
- primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint

<a id="canonical-2100332123103230-2213333330233032-1113231100012320-0211332122323122-0010033302130021-1112322003202312-0013033222301123-1012300030101212"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for sha256 fingerprint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("fingerprint")}
```

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

Terraform syntax:

```terraform
sha256_fingerprint {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113010202303031-2231130131103210-3011311232302201-0133100202302302-0230111112120230-0101221032212310-3201201203002313-1302230202233000"></a>

## Direct properties — sha256_fingerprint / 013200211003 / 3

<a id="canonical-1203210033210132-2221003233232103-1230233003103323-3030131030010303-1010112211003312-0013220311321321-0112003033021310-1100203131022201"></a>

<a id="canonical-0133111222323000-1131002202132013-1111111133210333-0013112023020103-0233002112312233-1030332230201033-0010010223210030-2100200211301310"></a>

## fingerprint property — sha256_fingerprint / 013200211003 / 4

Type: `"string"`. Optional.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(64, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-0332032223320031-0000311301110232-1031311203101301-2013301130221300-3133231113032123-1130100201002012-1120102311222333-1310103032221222"></a>

## Next pages — sha256_fingerprint / 013200211003 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](resources--dns_zone--reference--group-003.md#canonical-1112321033032121-3002101020012323-0210321303013102-1231333231113030-1300221302123321-2023012102301030-2202201320321231-1022101102121322)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0031230210322331-3321210223301010-0132333332013133-0020032020210333-3021212331022022-3312031103113322-3133100300013111-0030030101213110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311322302310121-3001232133213123-0001303010123322-1131012223021003-1311310012231332-3233031312101003-3312213111220123-3303300113322030"></a>

## primary.rr_set_group.rr_set.tlsa_record — tlsa_record / 201013313003 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.tlsa_record

<a id="canonical-3310230100220110-2003230230011010-1023111321302213-2230131233321320-3001222102312131-3232032301310331-0320213001312302-2322332212011113"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tlsa record.

Upstream description:

DNS TLSA Record.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
tlsa_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001020323031130-0210201032123002-1200110120131130-3023112033302212-2211131221032033-3302103202000323-3122022132222300-1130333132201131"></a>

## Direct properties — tlsa_record / 201013313003 / 3

<a id="canonical-1311310012330022-1003131133330232-1230203212321112-0012021010333112-2320101211130013-0012211102012103-0022200321311220-0230020023123011"></a>

<a id="canonical-1002133133030220-3032021133311222-1131212023003313-3303220213221313-0301022300100231-1332201202123020-1303001330131201-2200322020033321"></a>

## name property — tlsa_record / 201013313003 / 4

Type: `"string"`. Optional.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](resources--dns_zone--reference--group-003.md#canonical-3133001001210321-3322222312020213-2011103303030121-0113113320130302-0332233032210232-0233210203113001-0330002202013030-3223320001022020): complete subsection reference.

<a id="canonical-3022131210031312-0021123121131311-1323112212131110-1101202311211222-1113313012331203-1333130130222301-1203023001222011-0123113330220330"></a>

## Next pages — tlsa_record / 201013313003 / 5

- [primary.rr_set_group.rr_set.tlsa_record.values](resources--dns_zone--reference--group-003.md#canonical-3133001001210321-3322222312020213-2011103303030121-0113113320130302-0332233032210232-0233210203113001-0330002202013030-3223320001022020)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3133001001210321-3322222312020213-2011103303030121-0113113320130302-0332233032210232-0233210203113001-0330002202013030-3223320001022020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210302230302221-0122220321103230-3112130111023123-1111133131303332-3210213221202120-3222310133233011-3111033101221320-0212122221330333"></a>

## primary.rr_set_group.rr_set.tlsa_record.values — values / 121121322211 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0031230210322331-3321210223301010-0132333332013133-0020032020210333-3021212331022022-3312031103113322-3133100300013111-0030030101213110)
- primary.rr_set_group.rr_set.tlsa_record.values

<a id="canonical-2211233033003132-3231130220321230-1111303101003111-3321000232232120-1212221101223233-3320200121020020-1232201123102121-0333003133110300"></a>

Type: `"object"`. list nested block, Optional.

TLSA Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_association_data")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232120010000322-0121100102100233-3231222201113310-3310233221121301-3233232201301303-3230330022013011-2133011320032201-0010020002200312"></a>

## Direct properties — values / 121121322211 / 3

<a id="canonical-2011321131202002-2100102112300020-3330121221020233-2132110032030031-3202131212312122-0100100011213313-3121100203120303-2020201311021230"></a>

<a id="canonical-3200332022320111-0210302232331112-1002332223131030-2231231321120032-2201032233033202-1211220300001301-3221012033322022-2322103213030031"></a>

## certificate_association_data property — values / 121121322211 / 4

Type: `"string"`. Optional.

The actual data to be matched given the settings of the other fields.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1011330200330030-2123312103302022-2113100323023323-1302211300022100-1130312210331310-1103233313330103-3322203320322103-2330211330322212"></a>

<a id="canonical-3100132233202121-3120001020212231-2200322100010313-1121222201002323-1232300231211321-2021230310101121-1203102202113213-3000132113121011"></a>

## certificate_usage property — values / 121121322211 / 5

Type: `"string"`. Optional.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Upstream description:

&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint

&#8203;- ServiceCertificateConstraint: Service Certificate Constraint

&#8203;- TrustAnchorAssertion: Trust Anchor Assertion

&#8203;- DomainIssuedCertificate: Domain Issued Certificate.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["CertificateAuthorityConstraint","DomainIssuedCertificate","ServiceCertificateConstraint","TrustAnchorAssertion"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120103313311223-3311212033320132-2313021132103012-3120322113032010-2210102330331300-1231213031030011-2233131013311220-2333023230333201"></a>

<a id="canonical-2213210120011221-2023203221013122-3301012303100123-2132103012300001-3200320222113000-2203230103122221-3302203021213331-0020111113322333"></a>

## matching_type property — values / 121121322211 / 6

Type: `"string"`. Optional.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Upstream description:

&#8203;- NoHash: No Hash

&#8203;- SHA256: SHA-256

&#8203;- SHA512: SHA-512.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["NoHash","SHA256","SHA512"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NoHash",
    "SHA256",
    "SHA512"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1231000011230302-1002130201101212-3023110333020223-3030001121300001-0200012003331321-2103210113120121-0311330322323123-3221133300203330"></a>

<a id="canonical-2112030312312102-2301231213311001-1230221202320230-1221030300120222-0201020230021120-0332030102222031-0211011021013201-3203310311233200"></a>

## selector property — values / 121121322211 / 7

Type: `"string"`. Optional.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Upstream description:

&#8203;- FullCertificate: Full Certificate

&#8203;- UseSubjectPublicKey: Use Subject Public Key.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["FullCertificate","UseSubjectPublicKey"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("FullCertificate",
    "UseSubjectPublicKey"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0212131020002113-1032013022032321-1203100202002320-1203120212103220-2323000100033213-3310113001132100-1132010330122031-2003232020203203"></a>

## Next pages — values / 121121322211 / 8

- [primary.rr_set_group.rr_set.tlsa_record](resources--dns_zone--reference--group-003.md#canonical-0031230210322331-3321210223301010-0132333332013133-0020032020210333-3021212331022022-3312031103113322-3133100300013111-0030030101213110)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1121131111220220-0333230311100110-3220333331113330-3330020320131230-1033022131302013-0333232321103112-3113013123231300-0230232233102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221223201330311-0103102232111101-3030301301313302-3010221031001133-1031313203311020-0211002331010313-1021003111303333-2121221230113013"></a>

## primary.rr_set_group.rr_set.txt_record — txt_record / 012311210010 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [primary.rr_set_group](resources--dns_zone--reference--group-002.md#canonical-1231233131102130-3033122123213221-0200212121031323-0310313200103100-1133203333000312-0022023033122202-1122010230320231-0230033302322110)
- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- primary.rr_set_group.rr_set.txt_record

<a id="canonical-1100331110322030-0201111130213202-3110023110000301-0003213100033201-1112001001223102-1320003200303132-3000323323020202-0033331320301332"></a>

Type: `"object"`. single nested block, Optional.

DNSTXTResourceRecord.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("values")}
```

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

Terraform syntax:

```terraform
txt_record {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111010303301021-0100333022233020-1132100130323201-1022012100003032-2030311303302223-1210020223313021-2132112223332102-3220300100231030"></a>

## Direct properties — txt_record / 012311210010 / 3

<a id="canonical-2030201221211031-0012002020130111-1310002001030002-0230223203110013-2302201111203001-2113210032213321-2332032321221200-2223333202010210"></a>

<a id="canonical-2211221311001302-1321011131132321-2320203231012221-0130121121030212-2012122222000022-1232201121101303-3203203103013002-3010221120211201"></a>

## name property — txt_record / 012311210010 / 4

Type: `"string"`. Optional.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-1213202100110122-1000332210102110-2233200221331210-3112130123122210-3131301222211012-2011200123210122-0203022103313330-2032101302022210"></a>

<a id="canonical-1100122232002123-2332200003203322-0231110311013210-3313133202322312-3313000023022222-1203113102312223-0123201232223323-2221010232120313"></a>

## values property — txt_record / 012311210010 / 5

Type: `["list", "string"]`. Optional.

Text. Configuration parameter for values

Upstream description:

Configuration parameter for values

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1300130103200113-3221133302132132-0113333031031202-0111112033203031-3120311233200211-2332120132233003-0231112302212112-2023033022021003"></a>

## Next pages — txt_record / 012311210010 / 6

- [primary.rr_set_group.rr_set](resources--dns_zone--reference--group-002.md#canonical-3302221210003330-0233000030131300-2222332320003010-3202211132313210-1200111232120201-2121200231313301-3123021322110010-3031200330020313)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1020202022030312-2130233001011230-0123121101212023-0220103223030203-2021100102202011-2332311013003301-3000013203013302-0121111103000130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211033221331230-2112210011303023-1023321313012233-3233331001103133-0222020033300023-2210201130323203-0223013311213301-0220202213021103"></a>

## primary.soa_parameters — soa_parameters / 311233131001 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- primary.soa_parameters

<a id="canonical-2121220321110032-3211203313030203-0103020311122201-0023320210202020-3121000202111212-1213011301133030-3111130103212201-0011312132222212"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for soa parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("refresh",
    "retry")}
```

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

Terraform syntax:

```terraform
soa_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221210121321221-3323221123013020-2131000200200021-0022320023001120-1132012221200200-1002031102300101-0221003110203310-3030102322210223"></a>

## Direct properties — soa_parameters / 311233131001 / 3

<a id="canonical-3230132300000033-3110101120232103-0102232131022221-0213122210211302-1321012320222201-0122020231231331-3331220311231102-3102322313020311"></a>

<a id="canonical-2202110132121202-0333303330312000-2012001033122232-3000000121300111-2001032203303100-0121330030212000-3211033210103032-0023312001011001"></a>

## expire property — soa_parameters / 311233131001 / 4

Type: `"number"`. Optional.

Expire value indicates when secondary nameservers should stop answering request for this zone if
primary does not respond.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-2321101100133013-2133313330023311-0230331332003001-3012220301332222-0102123320302303-1312030230112333-3112320201003201-1330311010030021"></a>

<a id="canonical-1030301223021301-2201200022322003-3000100030132213-0121002332112301-1302220130110230-2300302232112230-2331311032301333-1320122223332023"></a>

## negative_ttl property — soa_parameters / 311233131001 / 5

Type: `"number"`. Optional.

Negative TTL value indicates how long to cache non-existent resource record for this zone.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-2312322201230120-2102133302301003-1020333031120112-3232213210311122-3322000112101202-3011312002030010-3001102333031203-2200333102220312"></a>

<a id="canonical-3301232330013020-2011231320002332-2311200201012223-2112130110212333-0003122211002210-2333311122212221-3022313212012332-2003313231021031"></a>

## refresh property — soa_parameters / 311233131001 / 6

Type: `"number"`. Optional.

Refresh value indicates when secondary nameservers should query for the SOA record to detect zone
changes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(3600, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-0132223313310300-3002321121310323-1322333201102312-3002131301301323-2310010230213030-1221131321220023-0222013132031032-0201101302133211"></a>

<a id="canonical-0202221320332100-1113132120002311-3302323130131113-1131201022332110-3212300102012103-0110032322230230-1103013221312233-2030033113300003"></a>

## retry property — soa_parameters / 311233131001 / 7

Type: `"number"`. Optional.

Retry value indicates when secondary nameservers should retry to request the serial number if
primary does not respond.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(60, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3031121330002222-1122010031033213-2002203113122001-2220320201322012-3303123323323323-0312201022120333-3312123322122323-2113131330202000"></a>

<a id="canonical-2213030132102121-2022303112130100-3120121302203101-2322203101211012-1013002202222222-0013002113323301-0012330131113333-2120231130001112"></a>

## TTL property — soa_parameters / 311233131001 / 8

Type: `"number"`. Optional.

TTL. SOA record time to live (in seconds)

Upstream description:

SOA record time to live (in seconds)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3233032223032330-3301211312033023-0222120002333022-3322331313122222-3013032321122022-2132310333013201-2021031200012133-1311221203111301"></a>

## Next pages — soa_parameters / 311233131001 / 9

- [primary](resources--dns_zone--reference--group-001.md#canonical-3123202203020311-2202302013313012-1330131221123210-3030331200201223-0320013021212223-0000332003102102-1303330133001310-3322331220312312)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222013133003013-1021032013210130-1110302021230000-2112222133102000-3131023110210030-1101203003011232-1303222202332201-2020102021010000"></a>

## secondary — secondary / 230020032002 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- secondary

<a id="canonical-0223010131102331-0310121320000300-3131133100321303-1312200031031101-2101222213223300-3110301303132223-1023132100303123-3012223223200121"></a>

Type: `"object"`. single nested block, Optional.

SecondaryDNSCreateSpecType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_servers")}
```

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

Terraform syntax:

```terraform
secondary {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320121013310121-0010123111303231-0030000231021102-1213221331320010-2112311110310012-3333333213022031-2302001033303322-0230120230011003"></a>

## Direct properties — secondary / 230020032002 / 3

<a id="canonical-3022322020202002-2320120232202130-1332110110120132-3231131222233013-3023023031131313-2000321032320123-0222022323233201-0223032300003022"></a>

<a id="canonical-2003000301233323-0123110111021213-1301133331333321-3002100130301123-2100213103310132-0022121103222110-1100131031332213-3321322302322122"></a>

## primary_servers property — secondary / 230020032002 / 4

Type: `["list", "string"]`. Optional.

Configuration parameter for primary servers.

Upstream description:

Configuration parameter for primary servers

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 10),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131221012001030-0213220100322000-0121123330330200-3122302010233231-3322322030130132-1232101300233203-2022220033103233-2112112211323031"></a>

<a id="canonical-3200131202312000-1330132102111331-0022133202020232-3103103002201321-0302302132303223-3302022230000023-0311012313100101-2211011312201213"></a>

## tsig_key_algorithm property — secondary / 230020032002 / 5

Type: `"string"`. Optional.

\[Enum: HMAC\_MD5|UNDEFINED|HMAC\_SHA1|HMAC\_SHA224|HMAC\_SHA256|HMAC\_SHA384|HMAC\_SHA512\] TSIG
key-value must be compatible with the specified algorithm - UNDEFINED: UNDEFINED - HMAC\_MD5:
HMAC\_MD5 - HMAC\_SHA1: HMAC\_SHA1 - HMAC\_SHA224: HMAC\_SHA224 - HMAC\_SHA256: HMAC\_SHA256 -
HMAC\_SHA384: HMAC\_SHA384 - HMAC\_SHA512: HMAC\_SHA512. Possible values are \`HMAC\_MD5\`,
\`UNDEFINED\`, \`HMAC\_SHA1\`, \`HMAC\_SHA224\`, \`HMAC\_SHA256\`, \`HMAC\_SHA384\`,
\`HMAC\_SHA512\`. Defaults to \`UNDEFINED\`.

Upstream description:

TSIG key-value must be compatible with the specified algorithm

&#8203;- UNDEFINED: UNDEFINED

&#8203;- HMAC\_MD5: HMAC\_MD5

&#8203;- HMAC\_SHA1: HMAC\_SHA1

&#8203;- HMAC\_SHA224: HMAC\_SHA224

&#8203;- HMAC\_SHA256: HMAC\_SHA256

&#8203;- HMAC\_SHA384: HMAC\_SHA384

&#8203;- HMAC\_SHA512: HMAC\_SHA512.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["HMAC_MD5","HMAC_SHA1","HMAC_SHA224","HMAC_SHA256","HMAC_SHA384","HMAC_SHA512","UNDEFINED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UNDEFINED",
  "enum": [
    "HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1312113303301123-0131313203312122-2000203020002221-2012011313232100-0302313003131321-3031110232312212-3300111032232002-3020122021101320"></a>

<a id="canonical-1103011020101310-0212102130222230-2232321302120223-3032303330002111-1110333120201202-2022222132103311-2131213113200032-2213033030203320"></a>

## tsig_key_name property — secondary / 230020032002 / 6

Type: `"string"`. Optional.

TSIG key name as used in TSIG protocol extension.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013): complete subsection reference.

<a id="canonical-3313123201013120-0232011120210002-1000201332313002-3031311020002220-3022231313303331-2120112020123211-2212333233033331-3223331101300020"></a>

## Next pages — secondary / 230020032002 / 7

- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320033122100010-3011202300022113-1031220231312233-0201031130303002-0003230202311002-1012312313121132-2023130000221003-2121010000330310"></a>

## secondary.tsig_key_value — tsig_key_value / 010020213233 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021)
- secondary.tsig_key_value

<a id="canonical-1133331300221133-0123021301132223-1331302203102230-0220012222331103-1201221303001002-3103010003311200-3002102213203003-3012013312111321"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
tsig_key_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221121333123130-3202010031323130-2232303001022303-2001132003020000-3112011111221033-2011130331312120-3312221312000312-3132123310103123"></a>

## Direct properties — tsig_key_value / 010020213233 / 3

- [blindfold_secret_info](resources--dns_zone--reference--group-003.md#canonical-0021202031221212-3320231031003200-0031123320023132-3131020332100030-1003112133001300-1112332332210033-1323202102222310-2123013130232330): complete subsection reference.

- [clear_secret_info](resources--dns_zone--reference--group-003.md#canonical-3302323322210222-2331232232101310-2233022013321120-2000233221232030-3211302223220000-2103131213212123-2121333200012003-2133032302002300): complete subsection reference.

<a id="canonical-3331301021011330-2203013002113220-2131210012223313-2002131222203210-0330300003122233-1222022222113032-2320322303010112-0321232112331133"></a>

## Next pages — tsig_key_value / 010020213233 / 4

- [secondary.tsig_key_value.blindfold_secret_info](resources--dns_zone--reference--group-003.md#canonical-0021202031221212-3320231031003200-0031123320023132-3131020332100030-1003112133001300-1112332332210033-1323202102222310-2123013130232330)
- [secondary.tsig_key_value.clear_secret_info](resources--dns_zone--reference--group-003.md#canonical-3302323322210222-2331232232101310-2233022013321120-2000233221232030-3211302223220000-2103131213212123-2121333200012003-2133032302002300)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-0021202031221212-3320231031003200-0031123320023132-3131020332100030-1003112133001300-1112332332210033-1323202102222310-2123013130232330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310112332030222-1120111103011033-2021021013321031-3310202102213112-2001223120300023-3202301300211122-1213020301330232-1330302130000011"></a>

## secondary.tsig_key_value.blindfold_secret_info — blindfold_secret_info / 032211300212 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021)
- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013)
- secondary.tsig_key_value.blindfold_secret_info

<a id="canonical-3301012122133333-0110202122221123-0220023330200102-2000211100012300-0130332230302202-1322101030213033-2113110313221230-1332020312120123"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233232233123313-1111322232221213-2211231132310331-1102301331323030-3103102233330010-1211211302133101-1311021022110012-0013301333022230"></a>

## Direct properties — blindfold_secret_info / 032211300212 / 3

<a id="canonical-2033213322032312-0311201200230231-2001101223023201-1330222303203323-0301121110300103-1313212221133112-0100323011301232-3132213223101233"></a>

<a id="canonical-0133013332322201-3310112100231232-3100330202102301-1133220232012031-1000203323011313-2320221130210033-3201120023333220-0321320210213333"></a>

## decryption_provider property — blindfold_secret_info / 032211300212 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3023330002313330-0310233121031302-3013213121310201-3103011010012001-1331201301011322-0131330133132222-2031020230021022-1303333022323331"></a>

<a id="canonical-0102200013221321-0002103330202103-1313032333323100-3232002001110002-2330213102022201-3110303113232211-1130030203223203-1121332311332221"></a>

## location property — blindfold_secret_info / 032211300212 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0231213330033102-3003110201202211-0202013213022303-0023000031003200-2101223332313210-0102312301311012-0013003223003233-0112100130310201"></a>

<a id="canonical-0321120102012211-1223130231231330-2001020102320131-0202311110100303-3210002212233132-3231302213003033-3310301111102301-0322202123220033"></a>

## store_provider property — blindfold_secret_info / 032211300212 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1001213201222111-2333221303130100-1033230221110122-1330112010333020-1102102123000313-3133303313330013-0030302113010230-2101133223013130"></a>

## Next pages — blindfold_secret_info / 032211300212 / 7

- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-3302323322210222-2331232232101310-2233022013321120-2000233221232030-3211302223220000-2103131213212123-2121333200012003-2133032302002300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321112001020332-1010001233221131-2332233233132002-1301222113313311-1130221222302330-0231010220302101-3131321121021301-1020300132102333"></a>

## secondary.tsig_key_value.clear_secret_info — clear_secret_info / 203101233012 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [secondary](resources--dns_zone--reference--group-003.md#canonical-2101113233122202-1230223131033031-2301310003322033-0133201323022132-1232110221020011-1332203122201231-3233232102003032-1103322203113021)
- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013)
- secondary.tsig_key_value.clear_secret_info

<a id="canonical-2023132111012003-0311221323013101-1122000230223032-2103123200220102-1023020011231121-3211122310021031-3232230201032312-0013130200310230"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132110110112221-1011020333010120-0312213330311223-0001112030120131-0012033313023102-2120100103133223-0122003231123301-3102320132211313"></a>

## Direct properties — clear_secret_info / 203101233012 / 3

<a id="canonical-0313132220201333-2033222033331013-0030330303311112-0322211020220301-0102331021000110-3213030222322230-3021000030321003-0103110133023013"></a>

<a id="canonical-2001123003230303-2102312022123112-2330212111033032-3010310310210222-0032210212131130-0130121233103322-2000022213132010-1223020123020102"></a>

## provider_ref property — clear_secret_info / 203101233012 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2222023222322022-1101121232133022-3110222320132113-3212123330200313-3031101201201313-0210120111300132-2003201020113203-2113012332131101"></a>

<a id="canonical-0321002230100122-2201033302313122-1033011231302032-0003020101233030-3113330020321210-1003113320220301-1110201312321112-1130320201100311"></a>

## URL property — clear_secret_info / 203101233012 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3203021121301212-2111311203221101-2330321312102321-3322323103123113-3110130300110311-0100232212033120-0131221320311203-3321030320310231"></a>

## Next pages — clear_secret_info / 203101233012 / 6

- [secondary.tsig_key_value](resources--dns_zone--reference--group-003.md#canonical-2000123211322222-0003033302032232-2123210100321132-1103211212210021-1032131231201123-1211232110111012-0303312322103132-2310001302330013)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)

<a id="canonical-1121023120323203-0300301033303311-3210213210020103-1313032000222330-2030312122211110-2210030103233231-0231221302231331-3010100002112333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213302201023132-3320101001313231-1030133033302300-0303133313122032-3123312232032003-3220113030212102-1103211310223232-1332212013313002"></a>

## timeouts — timeouts / 233331022221 / 2

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- timeouts

<a id="canonical-0300320321333023-0310332010031331-3031222113133311-3100030302022003-3212120202021302-0032200302202012-1100200031002133-3222120010110101"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020022210333132-0301020230022032-1101103330231130-1332003021112331-1301320000330003-1301122223010130-0332121130331111-1232001223323301"></a>

## Direct properties — timeouts / 233331022221 / 3

<a id="canonical-0300131231131222-2103011332202020-0301112203313120-0211010303001323-2310331101022332-1211200311113120-0121002013210222-0231323222020212"></a>

<a id="canonical-3323013313320030-3130123230322232-3002132200311121-2100132010303210-0022212313001011-1211213321032123-3320022320312200-3121210222010323"></a>

## create property — timeouts / 233331022221 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3321101022212312-1233211320310103-0013133032000303-0231010331321222-0213012223111311-0323213230221330-1133220321110110-2203021323101221"></a>

<a id="canonical-0332220011101031-3202002103220220-1202131101030223-2000102313020210-3332102222210230-0313020302120012-0133100300010010-0032201320102010"></a>

## delete property — timeouts / 233331022221 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0321030200103001-1021031232022232-1030103103022221-2103022311101312-3111010301003211-1221312222231132-3121023010300312-3023202012113020"></a>

<a id="canonical-3131331210331322-0112223120000310-2330321010023100-2033330230313101-0101211201112212-3131332221121000-3123213020032112-3232030032130332"></a>

## read property — timeouts / 233331022221 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3323010230210303-0302232133303033-0311231221113031-1231331231312321-3203133220303021-3331110110321332-3232112303121310-3220332113032321"></a>

<a id="canonical-3231003121023230-1301233332310022-0000100303211201-2022132133330333-2301023220001300-2020121000112000-1223112332212003-1200000110033323"></a>

## update property — timeouts / 233331022221 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0120131332202332-0000230200001330-2332031131320000-0003020113003200-2232010331323111-1233313100003012-1021002213021221-2203012002003113"></a>

## Next pages — timeouts / 233331022221 / 8

- [Property reference](resources--dns_zone--reference--group-001.md#canonical-0002310333001131-3221310230020300-1100323210113110-3011332112013233-0233233022020303-3111100323111003-0322213130311210-0302210132212031)
- [xcsh_dns_zone](../resources/dns_zone.md#canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323)
