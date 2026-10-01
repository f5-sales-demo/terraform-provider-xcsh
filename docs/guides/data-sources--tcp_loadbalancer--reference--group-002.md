---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-1320131302312321-1020002131003220-2310020303313312-2011203311202202-1302333203001210-0223001012020003-0101300320132123-3332312022213203"></a>

## advertise_on_public — advertise_on_public / 323013022131 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- advertise_on_public

<a id="canonical-2200031202211021-0313200202312233-1013023211320030-3123000101033002-1113110302231111-2101321210002322-3331003310213213-0312330021311122"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-0123021201233100-1201332103021123-2231010212013212-2123300323122111-1322211322022212-1231110221013231-3222031332012113-1331012222332031"></a>

## Direct properties — advertise_on_public / 323013022131 / 3

- [public_ip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2323110201301120-0232100210311033-3213000230012232-0121301023200133-3003210021330131-0331020233211323-0020203312031031-0023232033013202): complete subsection reference.

<a id="canonical-2202301110032312-1101231300302302-3123212203302023-3000032233110130-1232301100033023-0312311001121321-1231321313323133-0111123321200021"></a>

## Next pages — advertise_on_public / 323013022131 / 4

- [advertise_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2323110201301120-0232100210311033-3213000230012232-0121301023200133-3003210021330131-0331020233211323-0020203312031031-0023232033013202)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2323110201301120-0232100210311033-3213000230012232-0121301023200133-3003210021330131-0331020233211323-0020203312031031-0023232033013202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233212233021022-1213112100213022-2323111230203103-3223033332301212-1221132200220301-1020212220312112-3102322201131101-2101032123311112"></a>

## advertise_on_public.public_ip — public_ip / 203332301223 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1112201032021131-3221000132200333-3131001223100003-2232313233132313-2111023302121223-0310220211131233-2130013222232323-2223231123010031)
- advertise_on_public.public_ip

<a id="canonical-2302330022113120-0132220223113030-1203002312332220-3011132222110320-0100103320332132-3231213110010122-0221301003313120-3312121011200103"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1203020111330302-1123103320003330-1233333332132320-2221110120200322-0213022312111222-0211230201222232-0022303113330203-2200132033332323"></a>

## Direct properties — public_ip / 203332301223 / 3

<a id="canonical-1222030022130330-0331013312310213-0220003120000133-2122222123123030-1223301100031132-0103131211033302-3113031021023121-3022122220220013"></a>

<a id="canonical-0330321023210313-0113021023023330-2232233300103101-1323103110121022-0320332230100013-2220200111222222-2123321130031101-2302021210211022"></a>

## name property — public_ip / 203332301223 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0232110030303000-3323233312210002-0331002220113231-1222032131332110-0320133301131131-0232310110113333-1331330112132330-2303112220222012"></a>

<a id="canonical-0132323132203011-0101102121023013-2110101313031121-1010130031210323-3201130221103313-1001210121303321-0300232130313011-3111323011331031"></a>

## namespace property — public_ip / 203332301223 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2323212133012203-2231313332303232-0213313201013313-2200010311322211-1221211112100130-0112221313113110-0213322230323013-3310210023331003"></a>

<a id="canonical-0211310230022231-2103332320112220-3021010233212222-3030033010211310-3123323033313323-2222200020133320-0110022203312200-3301110210131210"></a>

## tenant property — public_ip / 203332301223 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1301003320230212-2001230133023132-2321313123002110-1033312201322331-0132230200132303-2021000021030031-0220330333112022-3121213213122113"></a>

## Next pages — public_ip / 203332301223 / 7

- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1112201032021131-3221000132200333-3131001223100003-2232313233132313-2111023302121223-0310220211131233-2130013222232323-2223231123010031)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2103301002331000-0011121122110023-2213202121233302-1122011020100033-1201330300100321-3021121333233123-2002013133223103-3102121301033313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033131030321320-2231223013130213-0221001130102331-1203101020103002-0101310100300132-0000122122311003-0223133000221111-2232210013011230"></a>

## advertise_on_public_default_vip — advertise_on_public_default_vip / 202331321302 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- advertise_on_public_default_vip

<a id="canonical-1123330002330211-1310302201203201-1113023310221132-0301203233322302-0032103302032320-2010301030313101-2213001103002101-3232220103002213"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3222031131232211-3112320133232130-2130101032333310-3103011013320120-3001303001100022-1322002231303013-1022033213120313-0112102212032003"></a>

## Direct properties — advertise_on_public_default_vip / 202331321302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313302100002211-1200132101011230-2101112303320001-2222121303103011-2212132202201321-3113100223231212-3000012232330110-0300301113102233"></a>

## Next pages — advertise_on_public_default_vip / 202331321302 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1120310112113021-0011220322222102-0130122230032031-0102010320303202-0001311211211213-3321333330113130-2001230311133133-1322321301312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001232013313300-3331233132321111-2122330203302002-1011331321121223-0303023130110323-0223200220313333-0231120013112233-0300300322133003"></a>

## default_lb_with_sni — default_lb_with_sni / 123310033321 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- default_lb_with_sni

<a id="canonical-3323001220100213-0332200120333110-0322222321133332-1232031323010000-3212121100010232-0301020212002130-0303222103123032-1320010121121321"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_lb\_with\_sni, no\_sni, sni; Default: default\_lb\_with\_sni\] Configuration
parameter for default lb with sni.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [default_lb_with_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3323001220100213-0332200120333110-0322222321133332-1232031323010000-3212121100010232-0301020212002130-0303222103123032-1320010121121321)
- [no_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2112113331223322-3231232001320233-1302030020311232-3100000332022230-1032022022211232-2010330311022330-0123212123130312-3120230231130212)
- [sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2312220212030333-0230231220111231-0012033022321102-0311002130002223-2112002002110030-0212103020302010-2021321302012013-2221101100130002)

Select alternatives according to the provider validators above.

<a id="canonical-0212102322130301-0222231203322302-0203120023331231-1313200030101031-3011102001010331-0032002321031323-0031130222223213-3023310310220322"></a>

## Direct properties — default_lb_with_sni / 123310033321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203111211300330-0301132222222132-3110222113331311-0212103303311001-0003333103323033-2133210112002012-3001103031001211-0221332231331320"></a>

## Next pages — default_lb_with_sni / 123310033321 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2301200001232003-2230220223123022-1031331131120332-3030120311113103-0010103000333222-3012312213013120-0313333310330302-1322122222220023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330301233021310-2110202313200321-2110111212120321-0211321023122020-2023302023022200-1020230321020211-1110023213230023-1013202320321320"></a>

## do_not_advertise — do_not_advertise / 320010013212 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- do_not_advertise

<a id="canonical-0001303000210302-2210122211320302-0002200332110310-0310113320302002-1122011331112123-3223311220302121-3230321301031211-3123012131133000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3221133123003120-3310101221330033-2313012010101000-0021333021300330-3211031322122100-1301232102020110-0001023033331033-2112113331332210"></a>

## Direct properties — do_not_advertise / 320010013212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130030312300100-3122333021121233-3000103022302030-2033231200033120-1302123002230100-1030211323332020-3322022031312130-2203310123333012"></a>

## Next pages — do_not_advertise / 320010013212 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2102221023203330-2101321113012132-3212312223313302-2333003020300013-3212302122333112-1203323112212131-0113122301023121-1123321312122331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232203033313001-0223203223330300-2113202312313302-0202020333122013-3310330132203123-1111212011230033-3321012301301131-0013310323312130"></a>

## do_not_retract_cluster — do_not_retract_cluster / 213320120122 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- do_not_retract_cluster

<a id="canonical-0021013132003110-0210102220132311-3302202220213101-2313122133232322-0211013312030111-2133331021111111-1102303210112013-1030322030021232"></a>

Type: `["object", {}]`. Computed.

\[OneOf: do\_not\_retract\_cluster, retract\_cluster\] Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [do_not_retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0021013132003110-0210102220132311-3302202220213101-2313122133232322-0211013312030111-2133331021111111-1102303210112013-1030322030021232)
- [retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1131023233233002-0323313303102032-2022112113312010-0331130133000123-0003323331110103-0100213211012000-1331100102121003-1302332022231233)

Select alternatives according to the provider validators above.

<a id="canonical-1303120100211321-2323022310023011-0201132232331303-2310112333321300-2102110031302221-1331033131212233-2103321200200003-0332212200312222"></a>

## Direct properties — do_not_retract_cluster / 213320120122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310030120102203-2230201230332033-1311113331220000-2022313133001203-3203311310000200-2233313120103210-2310013000021202-1222120112022230"></a>

## Next pages — do_not_retract_cluster / 213320120122 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1233323330321012-0022220130111220-0001311203331000-0032101313230230-0003312301003011-3322313221113031-0000100011331113-2312101310330012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001203231201311-1321031211010321-2200211130301123-0123302203221100-2013101110120023-3232310233012311-0003030311221023-2123010013333131"></a>

## hash_policy_choice_least_active — hash_policy_choice_least_active / 332310110021 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- hash_policy_choice_least_active

<a id="canonical-0213123311230102-3223303113132302-0222311331313011-2233121031121311-1102323230101312-1223301003121202-1122203220021330-1112321202221300"></a>

Type: `["object", {}]`. Computed.

\[OneOf: hash\_policy\_choice\_least\_active, hash\_policy\_choice\_random,
hash\_policy\_choice\_round\_robin, hash\_policy\_choice\_source\_ip\_stickiness\] Enable this
option

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [hash_policy_choice_least_active](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0213123311230102-3223303113132302-0222311331313011-2233121031121311-1102323230101312-1223301003121202-1122203220021330-1112321202221300)
- [hash_policy_choice_random](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3212033231123032-1011312023323013-1011123300220023-1231033203130221-1201031100111300-1010230221333212-1313301103132100-1122023101232212)
- [hash_policy_choice_round_robin](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0013003300111331-2221131030212002-0133311311011203-2121231303120023-1320011201221133-2232301003113132-2321033323301211-3331202302033223)
- [hash_policy_choice_source_ip_stickiness](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0001313102231321-2102202120232131-3313220313103221-2312113230312002-3312032203302111-0233332302110333-1313210300232111-1311102213312222)

Select alternatives according to the provider validators above.

<a id="canonical-0023033213110221-2303030200020010-1200102310321022-2002230222031231-1000132003331223-2210321313021232-0311221303310301-3011123011202112"></a>

## Direct properties — hash_policy_choice_least_active / 332310110021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102030032120223-2013322011213212-0032301101002313-0333221103032232-3210333011232202-1023013021021212-2022012031131222-1300213210031033"></a>

## Next pages — hash_policy_choice_least_active / 332310110021 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1323103221023221-2313321022023032-0102100310112221-3013030013003012-3002031302202202-0020002001010103-2300123101332113-1301322121013310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233220031210321-2000210032032002-1232302033122221-2313000103323203-1220321030313233-3010221113332001-0302030232220303-0130020133330032"></a>

## hash_policy_choice_random — hash_policy_choice_random / 302213233011 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- hash_policy_choice_random

<a id="canonical-3212033231123032-1011312023323013-1011123300220023-1231033203130221-1201031100111300-1010230221333212-1313301103132100-1122023101232212"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for hash policy choice random.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0312222203002303-2021012310031110-3202011300302012-2232331210032221-0033130100130331-0032011103310322-2133201030120210-1122131122102111"></a>

## Direct properties — hash_policy_choice_random / 302213233011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102322011332012-1220100223312331-1332111330230222-3120203233310220-0231022332312233-2031231201323222-3100033223000322-1010321333100200"></a>

## Next pages — hash_policy_choice_random / 302213233011 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2001222213202013-2110122010013020-0222012122321312-2221120003230033-3112000110322320-0030000330021323-3103332101033100-3231323011021122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210132010000131-3303000010200003-2023111233333231-2112311000120031-3221001321201303-3212202332222221-0100013212203112-2031021210213002"></a>

## hash_policy_choice_round_robin — hash_policy_choice_round_robin / 100112300201 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- hash_policy_choice_round_robin

<a id="canonical-0013003300111331-2221131030212002-0133311311011203-2121231303120023-1320011201221133-2232301003113132-2321033323301211-3331202302033223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for hash policy choice round robin. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0213323012313331-1132010120021002-2202103211332111-1301133002133013-3202120223232323-0010201033002110-0212232033212131-1201301100231101"></a>

## Direct properties — hash_policy_choice_round_robin / 100112300201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121230011112222-1323221001002210-3232220000033011-2321112033023222-3000130120100221-3033202313110111-1332032322101111-3000132011121010"></a>

## Next pages — hash_policy_choice_round_robin / 100112300201 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2101213003031210-0111210032000102-1202211230323123-2123332011131030-3210312021210331-1330112233100331-2210212303011011-3120212322222111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122023321221020-1120013211013031-2003020311211111-0100031221222201-0131102000333121-3321002000221302-0300221330322302-2320333200100220"></a>

## hash_policy_choice_source_ip_stickiness — hash_policy_choice_source_ip_stickiness / 201210201320 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-0001313102231321-2102202120232131-3313220313103221-2312113230312002-3312032203302111-0233332302110333-1313210300232111-1311102213312222"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2111120211312133-1333010132332330-1230123030003301-0333131213022301-3212212111232213-3101232003132321-2332231222003100-2030233210230113"></a>

## Direct properties — hash_policy_choice_source_ip_stickiness / 201210201320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032220210213012-1123123000300020-1211112031032000-2330333020012123-0013122002110031-3120313233120020-1010013031331300-0011122011010132"></a>

## Next pages — hash_policy_choice_source_ip_stickiness / 201210201320 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3232303121303303-3310101031211332-3320310013303113-3130011003231312-0221113013330021-0103001323103000-3230002332010321-1123002233310003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100002101302112-0302033313112021-1122023103113102-1220123132132232-0302220113233100-2020320010013323-2102232110213202-1220300111212321"></a>

## no_service_policies — no_service_policies / 032110112222 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- no_service_policies

<a id="canonical-0121112133023001-2133103202133131-1030333112200300-0221311100103201-3330222002023012-0321002320113000-0310301333203033-2200011113112222"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no service policies.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1120220332021130-3033131102200013-0303111023221331-3200333221011131-3132300200000332-0313232131020032-0113120011232213-0300032223132320"></a>

## Direct properties — no_service_policies / 032110112222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001211311312230-0202303130130212-1322121231013121-0002031022111300-2023102003032000-3130230201100312-2133103122202012-1102103321322201"></a>

## Next pages — no_service_policies / 032110112222 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3323311322022213-3030031301102132-3222213210012013-2113301030000013-1230313033002102-0010100310021223-2002303301021210-3332001213302330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021133231212300-2112320002133122-3030121031331101-0222320022001333-3331010221233321-3321000330130001-2330231021023122-1020021232003221"></a>

## no_sni — no_sni / 122311331320 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- no_sni

<a id="canonical-2112113331223322-3231232001320233-1302030020311232-3100000332022230-1032022022211232-2010330311022330-0123212123130312-3120230231130212"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3003211202101022-0220333000311030-1323311320131310-2331032220133031-0130133012113312-1010123112221030-2212120212020231-0020231002133322"></a>

## Direct properties — no_sni / 122311331320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210333333200101-3303131033322011-3300030131010201-1203003003300031-2312310223013230-3201310323010213-1112132021310131-1333000320330023"></a>

## Next pages — no_sni / 122311331320 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033333231322200-2212101031310100-1020023320203111-1311300203101332-3112003112010030-3131122030111331-0302223332223233-2032003103021120"></a>

## origin_pools_weights — origin_pools_weights / 120200111033 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- origin_pools_weights

<a id="canonical-3200202332003310-0211023331000103-0121311213201021-2001122321011101-1130023223300221-1121311210313223-3022112123003300-0100230133313000"></a>

Type: `"list"`. Computed.

Origin pools and weights used for this load balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0212303220110033-0033203002123112-1123111031032033-2323013021320333-3220213103110033-1023011331320233-0331113002122230-3000320223213302"></a>

## Direct properties — origin_pools_weights / 120200111033 / 3

- [cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1100323120213031-0021203033221103-0232210303113332-1232122132332320-2022312033230033-3203211222032232-2203103200232310-0120121202303230): complete subsection reference.

- [endpoint_subsets](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3233020300102002-3212200311012103-3303223032200213-0321120133210212-1310320312202012-3223222313112032-1320233023030102-3010212231102021): complete subsection reference.

- [pool](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3113021102231311-3020330031331101-3221032221322031-1022322103230303-1130012222122212-3122202323303323-3113103301320101-3130030230221032): complete subsection reference.

<a id="canonical-0103301112020110-2133210101322130-3120321133113130-1213031012300313-1233031101032300-2023001121123303-0030103230013311-1223211120321312"></a>

<a id="canonical-2030131001230230-2312321032310332-2023312103331003-3232122300301322-2321312103201312-2132033023122313-3121321033331333-2110110123220223"></a>

## priority property — origin_pools_weights / 120200111033 / 4

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1231132003213230-2312112323103213-1303032131103311-1020022230023332-2001001012010113-3201030003313213-2101020123120220-0123312003101312"></a>

<a id="canonical-3320110213022110-3301113122100033-2110221330101303-0102103313332031-1312000002031223-3303331201103331-3012121131201232-3330201213031200"></a>

## weight property — origin_pools_weights / 120200111033 / 5

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233201211101212-3233100313202322-2023103013323310-0123021203033020-2323211312232002-0220032131100202-3133110231132322-0232320100321132"></a>

## Next pages — origin_pools_weights / 120200111033 / 6

- [origin_pools_weights.cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1100323120213031-0021203033221103-0232210303113332-1232122132332320-2022312033230033-3203211222032232-2203103200232310-0120121202303230)
- [origin_pools_weights.endpoint_subsets](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3233020300102002-3212200311012103-3303223032200213-0321120133210212-1310320312202012-3223222313112032-1320233023030102-3010212231102021)
- [origin_pools_weights.pool](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3113021102231311-3020330031331101-3221032221322031-1022322103230303-1130012222122212-3122202323303323-3113103301320101-3130030230221032)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1100323120213031-0021203033221103-0232210303113332-1232122132332320-2022312033230033-3203211222032232-2203103200232310-0120121202303230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103030032332202-0212212312100131-1012331133221330-2131122221110112-0031010012312113-1112131133031210-1302333320302312-2013031000202032"></a>

## origin_pools_weights.cluster — cluster / 103320003313 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101)
- origin_pools_weights.cluster

<a id="canonical-1113012100011030-2001322102301030-1131010011033030-1313003012021101-0110011203030111-0032230200000311-3003320330302213-1200101331100111"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2023021211333022-1221132201201211-3113130301210113-3132101011001223-3000120111221103-0120221202103223-0102321031112223-1200022103032300"></a>

## Direct properties — cluster / 103320003313 / 3

<a id="canonical-2212013331103221-3031332033320031-1210122323123331-1031313113310310-3100223311311321-3120113311131312-3112322132002012-3220331130032303"></a>

<a id="canonical-2201311012021312-0120330111301333-1302011102133022-0101311221030223-1131303210112322-2210203123032332-1320302232120221-1030123310210121"></a>

## name property — cluster / 103320003313 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3200332330131231-3103203300030030-3031323302132210-3132221121211320-1230030120120212-0003023213012210-1221001331011023-1200032332022033"></a>

<a id="canonical-1330111000222011-3332332113202331-3113011010030312-3331001330113121-0012000232023210-2223320113303010-3132103300210101-2312332122033212"></a>

## namespace property — cluster / 103320003313 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0300331231310020-2221010013102111-3022100132132022-3312032231320012-2101033320221132-1123310201312103-1221200033010122-2222310330220323"></a>

<a id="canonical-0123000323330323-3030113103300201-1323020102203033-0321232020013312-3332121030303012-0230230233120212-0010210123301002-0302221220212131"></a>

## tenant property — cluster / 103320003313 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3333121120330311-2010032310123012-3301033212213002-1000300113300130-0110030230210311-3202022132311101-0202003011001210-2012100323323022"></a>

## Next pages — cluster / 103320003313 / 7

- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3233020300102002-3212200311012103-3303223032200213-0321120133210212-1310320312202012-3223222313112032-1320233023030102-3010212231102021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100112022303030-2101100210102301-1203210030311110-2111212103032313-2010332110233101-0132221201311231-1221310211121031-3310010332202011"></a>

## origin_pools_weights.endpoint_subsets — endpoint_subsets / 112212212122 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101)
- origin_pools_weights.endpoint_subsets

<a id="canonical-3220123132322001-0221230110222113-2301001313220031-2003111011112222-0210003310111311-0231121311123233-3110203332231311-1100133102333121"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

<a id="canonical-3233000220010022-1331002330030033-1331213122320132-2033102133021130-3122200220122233-1221223121232313-2213301332132021-3132120301222212"></a>

## Direct properties — endpoint_subsets / 112212212122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133021010120001-2023233023120122-2322311021112313-2322302232302031-3011200231313100-1320121220312111-3312120001222230-1333213201211111"></a>

## Next pages — endpoint_subsets / 112212212122 / 4

- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3113021102231311-3020330031331101-3221032221322031-1022322103230303-1130012222122212-3122202323303323-3113103301320101-3130030230221032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120003031302120-2211013231221101-0121001101202030-2133322111230220-1012301120121331-3000100011303203-1103313022323122-0301310013232232"></a>

## origin_pools_weights.pool — pool / 130031001233 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101)
- origin_pools_weights.pool

<a id="canonical-1132213133202332-1322012003101323-3330023000200011-0331233133011332-2222300223122200-0331130322222123-2231333130011310-3012311322322233"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0023003102301233-2002101101120201-0102020033210030-1110232032011132-0312112031332321-0202032120001201-0200200203302003-2010123202121033"></a>

## Direct properties — pool / 130031001233 / 3

<a id="canonical-1202230130301301-0102231022221021-1132312202023331-0323033331101101-2200020223321123-2133213330222033-3113001133033031-1130322302230120"></a>

<a id="canonical-2200022121222103-3332222012203303-2220020202100031-0321000222202023-2102132100033231-2002223221331101-3203231301312103-1211103012133201"></a>

## name property — pool / 130031001233 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3233100010300122-3230221023120200-1113012021131032-0020202101121311-3100220330130000-2333032032122032-2033230320321232-0220233100122310"></a>

<a id="canonical-2212233320033333-0231100232203310-2100310011300013-3031232233212132-0100131230302023-2000003332110200-2223123103333211-3201133121030331"></a>

## namespace property — pool / 130031001233 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3321202331223033-2020301031110310-2212010110322331-2220302023003201-3012111002301110-0111013231221032-1311310333220110-1123230331332102"></a>

<a id="canonical-2202322203003200-3322131332121321-1132102202123210-2100223311312112-0100213312313101-1113311302311003-2223303223032011-3123002301012303"></a>

## tenant property — pool / 130031001233 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1132330211032101-0301321202121022-0311022131123333-0212231032201030-0310330020230222-0322222112102202-3232301031232313-1301332130033012"></a>

## Next pages — pool / 130031001233 / 7

- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3030013003230233-1020032322120203-1023112003112123-2102002023033311-2321113323313010-0110100232101332-0123011112033022-0022330033023211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211321302022332-2131131320032031-0220002302212003-1221201010221221-2033101030113101-2110301122031330-2223013213020230-3311230120002031"></a>

## retract_cluster — retract_cluster / 011021231320 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- retract_cluster

<a id="canonical-1131023233233002-0323313303102032-2022112113312010-0331130133000123-0003323331110103-0100213211012000-1331100102121003-1302332022231233"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0232113001032300-3312333200033200-2111021231121233-1103223201021001-1021110311110133-3002121111332200-1010302221103231-3320301323103230"></a>

## Direct properties — retract_cluster / 011021231320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322022330300312-3301000020302003-0020001120213330-1202110311230020-3303213213120211-0132223032202101-1332010121300032-0111200222320032"></a>

## Next pages — retract_cluster / 011021231320 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0123020233112322-0223121002333303-2100322322223000-3300212133321100-3320102320231010-0321033323023300-3233233123103010-2121331023133111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113223223001303-3003311020131123-0311111030130220-1101320323110112-0030031223320220-2130033303010023-1000233022321331-3300203133223023"></a>

## service_policies_from_namespace — service_policies_from_namespace / 031213000231 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- service_policies_from_namespace

<a id="canonical-3301033031210010-2323133031222031-3022302003130302-0332310330103322-1101123130223311-2303021212131110-1301210012202313-1202323210102000"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2312301112331013-0300313213021131-1123322013320321-0102020021110121-1132003233131220-0002231132122220-2012132333002331-3122122332210031"></a>

## Direct properties — service_policies_from_namespace / 031213000231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113231101101013-2313203230201031-0021222210130301-3313013333201222-1331200322210231-1110023033110001-0323331000121113-1001123303300330"></a>

## Next pages — service_policies_from_namespace / 031213000231 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2133300001211310-3203211112301011-3012211322222112-3322000131211112-0133200313302211-3311012013010002-0330310112023122-1200113222020223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210202013111201-2201231020021001-0311032113101133-2012133103322133-1000200122123213-1002010033113122-1020211130132103-0030221323112212"></a>

## sni — sni / 012310113220 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- sni

<a id="canonical-2312220212030333-0230231220111231-0012033022321102-0311002130002223-2112002002110030-0212103020302010-2021321302012013-2221101100130002"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2123011330103003-2020132010123232-0222202210321230-1101121123111310-1133101323210021-2001231112002022-0103032321021333-0320010320300110"></a>

## Direct properties — sni / 012310113220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010313311223120-2100320211010100-1221130211113003-2010131330033201-3231301132100133-0122211020232231-3300303000322232-1220230010110312"></a>

## Next pages — sni / 012310113220 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2032030011011133-2021231112123000-0010123000231211-1233113301302232-0220010312011312-1000210021110111-1203111330233211-3303302223213330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100103301211333-0202202031101111-0123130120022202-2332002213003112-2100132132001001-1320333002020121-3030110233302010-0232010123312222"></a>

## tcp — tcp / 022103303003 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- tcp

<a id="canonical-1233311120231210-2331220321101311-1220130030202313-0102210322110321-2301130120100010-1311032331121332-3330332030322100-0111113320312013"></a>

Type: `["object", {}]`. Computed.

\[OneOf: tcp, tls\_tcp, tls\_tcp\_auto\_cert\] Enable this option. Defaults to \`map\[\]\`. Server
applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1233311120231210-2331220321101311-1220130030202313-0102210322110321-2301130120100010-1311032331121332-3330332030322100-0111113320312013)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1212102131222122-0010230212231023-2032230231022002-3301321213331102-1011221101033321-2331201103200201-1332121102010012-1210120021232232)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3122010013013333-0010120031233203-2230121020320021-2303221102121132-3030013013032022-1321333312130213-1313113220130030-1002023233330000)

Select alternatives according to the provider validators above.

<a id="canonical-0203210000012000-2310332201320202-1030001311300132-2210110023010233-2223200310313103-1122332200122232-1313232310322011-2231312313232031"></a>

## Direct properties — tcp / 022103303003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000022123333330-1013023022311132-2020122210312322-2203111003103220-2220323122221002-2332110333113001-0101101013311102-2220112013203202"></a>

## Next pages — tcp / 022103303003 / 4

- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323312010220012-0002122220102200-2122001023210330-1320221312023021-3332313123113202-0210000123103313-1120320200323231-2022021303032322"></a>

## tls_tcp — tls_tcp / 103031231011 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- tls_tcp

<a id="canonical-1212102131222122-0010230212231023-2032230231022002-3301321213331102-1011221101033321-2331201103200201-1332121102010012-1210120021232232"></a>

Type: `"single"`. Computed.

Choice for selecting TLS over TCP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-3322223221201323-3220330012223201-2323000020111220-1013303012232103-2132121202222031-0321020003331300-2113220320213223-2002203230232103"></a>

## Direct properties — tls_tcp / 103031231011 / 3

- [tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123): complete subsection reference.

- [tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120): complete subsection reference.

<a id="canonical-2120300220321230-2312113313333003-2110312302320300-0012021133233212-3301323020133310-0311001012321022-0000033033230331-2200000301202333"></a>

## Next pages — tls_tcp / 103031231011 / 4

- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211103100111332-1011212200303220-1330321311303131-2023021003203003-1232023003122120-3101222133011113-1210020331111302-2312122102023330"></a>

## tls_tcp.tls_cert_params — tls_cert_params / 221322211202 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- tls_tcp.tls_cert_params

<a id="canonical-3132223110233100-1320012202202033-1132100030231233-0033102321001200-2232203103300203-1033211111333321-0331200012002110-0312303033331120"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-2111333101212021-1313310100022302-2101220003132220-3131231000331022-2332223313113023-1211003213310022-1112033133320211-2200313113300122"></a>

## Direct properties — tls_cert_params / 221322211202 / 3

- [certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1112320022302113-3133132103111312-2013212002203203-2320231003120030-1120232102313313-3012130013022023-1331300202221120-2020231321233022): complete subsection reference.

- [no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2012003003113330-0132112313013121-1023333102003321-1133031032330010-0131232013331213-3300333230002230-2221123012220330-3330312310233012): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232): complete subsection reference.

<a id="canonical-0332021200031232-2023303121201110-0232213110302303-1020021302031103-2323320231000320-1000003321110031-1222223323022220-1010031301110233"></a>

## Next pages — tls_cert_params / 221322211202 / 4

- [tls_tcp.tls_cert_params.certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1112320022302113-3133132103111312-2013212002203203-2320231003120030-1120232102313313-3012130013022023-1331300202221120-2020231321233022)
- [tls_tcp.tls_cert_params.no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2012003003113330-0132112313013121-1023333102003321-1133031032330010-0131232013331213-3300333230002230-2221123012220330-3330312310233012)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1112320022302113-3133132103111312-2013212002203203-2320231003120030-1120232102313313-3012130013022023-1331300202221120-2020231321233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101131121012211-1032021210231100-1022031230023103-2213220022110300-3002002102011200-1103321331010032-1303032303030021-3021231322030202"></a>

## tls_tcp.tls_cert_params.certificates — certificates / 301301303120 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- tls_tcp.tls_cert_params.certificates

<a id="canonical-2131031203113123-2120023231133210-2031321322203322-3020220130231101-3232323000013013-2130030221033213-2333302213132323-0211123023031222"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2210223223201000-1311030103102321-2232303121023231-1202012132020332-0221330312030121-3231310001313102-0121031303321023-0220003013122303"></a>

## Direct properties — certificates / 301301303120 / 3

<a id="canonical-0330213130033302-2311133212003232-1231011202220331-0201310012311002-1302323032110110-2011103103003101-1102100200021002-2331331000332311"></a>

<a id="canonical-3121131323021321-2030030233000310-1201223210330311-3233001132000201-0303312213213222-2130002023310030-1302323030113113-1032333000332230"></a>

## name property — certificates / 301301303120 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2012221313301311-3121010301222010-1303302113013130-0321122033330101-3323232033310200-3202100230201211-1230222131023332-3303222101121120"></a>

<a id="canonical-1112012330203003-0132223031203330-0001122013020113-2303130113010031-3130112300120103-1333131321022221-0322031310203322-1311322020103322"></a>

## namespace property — certificates / 301301303120 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2032333032103301-1310203201210310-2213221300031023-1333112212202311-0011122300022032-1113001303130332-0232330033123311-3123100102033320"></a>

<a id="canonical-2202302210020302-3211002230133220-0100200223020222-1333310033103310-3021203222232212-2010202331133233-2023220222033313-0122321030231233"></a>

## tenant property — certificates / 301301303120 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0132121002300130-3211310213233102-0121120103020233-0133201131132300-1313300002211332-3233311230221031-3021031113201122-1330101111021100"></a>

## Next pages — certificates / 301301303120 / 7

- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2012003003113330-0132112313013121-1023333102003321-1133031032330010-0131232013331213-3300333230002230-2221123012220330-3330312310233012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112331202201212-0022303331031310-0301013020330032-3311302023100313-3132333032112102-2020001110121221-1201110323333121-3222202203003121"></a>

## tls_tcp.tls_cert_params.no_mtls — no_mtls / 221330233311 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- tls_tcp.tls_cert_params.no_mtls

<a id="canonical-1131332332222222-2211033122302333-1310003133220203-1113330202203211-2202111030333333-3331211322133001-3130032120120221-3223110203103302"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0220330310110223-1011323312101000-3123232112102323-2230210202233300-0321322013132313-0302012312300320-3002113220112123-3132230101331231"></a>

## Direct properties — no_mtls / 221330233311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111323333321212-1201212031113303-3202222332032312-0112130222233003-0233220122032333-1312311222202101-2322101332323301-0002020132010003"></a>

## Next pages — no_mtls / 221330233311 / 4

- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031011122112132-2312231330130111-0030002033120101-3322332333311330-1022121012231213-3303021000312312-3221033211001000-0001022030313202"></a>

## tls_tcp.tls_cert_params.tls_config — tls_config / 333323122201 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- tls_tcp.tls_cert_params.tls_config

<a id="canonical-1121001013333021-1103303101321122-1231003331210113-2031010221010210-1313221103013223-1111313301230202-0232031132211130-3311300223310031"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-1100312132121321-3103233030313032-1301132031322301-1020322000130200-3110123310220113-3310031201301132-0021310300322103-2000131220323202"></a>

## Direct properties — tls_config / 333323122201 / 3

- [custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3311333121130022-2211220203231132-2000203330332021-0112200333013101-1022113001303332-1013233232001233-1221232220000320-2121232223212100): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0231221023332112-0013111300113021-1321302003220333-0312313301310013-3102001013231122-3231021230020303-2102021300103010-1121020000012121): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3013211012113110-1100321110200123-3223300302021332-1222102203330102-3010300131111231-1000132312212100-1220213120001221-2222031113301120): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2212002321101031-2100301010312210-0132001322000130-2120332323120010-2301113321223232-2223310031112030-1111100220031111-2121312120221200): complete subsection reference.

<a id="canonical-3010012133033003-2122132011331312-3103332212032313-2313320221011101-0030202312000100-2031302001103111-3103012123030120-3001311102101313"></a>

## Next pages — tls_config / 333323122201 / 4

- [tls_tcp.tls_cert_params.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3311333121130022-2211220203231132-2000203330332021-0112200333013101-1022113001303332-1013233232001233-1221232220000320-2121232223212100)
- [tls_tcp.tls_cert_params.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0231221023332112-0013111300113021-1321302003220333-0312313301310013-3102001013231122-3231021230020303-2102021300103010-1121020000012121)
- [tls_tcp.tls_cert_params.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3013211012113110-1100321110200123-3223300302021332-1222102203330102-3010300131111231-1000132312212100-1220213120001221-2222031113301120)
- [tls_tcp.tls_cert_params.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2212002321101031-2100301010312210-0132001322000130-2120332323120010-2301113321223232-2223310031112030-1111100220031111-2121312120221200)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3311333121130022-2211220203231132-2000203330332021-0112200333013101-1022113001303332-1013233232001233-1221232220000320-2121232223212100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311011231302223-1331022033030200-0202022201212022-3132213011121100-3221223213120131-2121332323033000-2301031210223130-0001220131032122"></a>

## tls_tcp.tls_cert_params.tls_config.custom_security — custom_security / 313200221313 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- tls_tcp.tls_cert_params.tls_config.custom_security

<a id="canonical-1200022311111031-1221101303010133-3200312012113323-1202310320303130-2211212320213322-1310213121233320-2130002103122320-2103013221013303"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-1130130231332111-0101221022011320-3130233210112213-0211111011101020-0233122332002230-2320203230023022-1201111230020332-2233020321200300"></a>

## Direct properties — custom_security / 313200221313 / 3

<a id="canonical-2312011321003020-2120011003222131-2230031122022211-1022020022033303-0222131301032332-1312210331323223-3000101212122300-0203013312220020"></a>

<a id="canonical-1201001111003301-3010123122302330-1300130000122231-1020100313012120-0210130211000220-2323330102132010-1210320332200133-1112310210013013"></a>

## cipher_suites property — custom_security / 313200221313 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2031121013232310-2212011011310221-1020112310311120-3300333311021221-2333120332310220-3111122010013033-2321113300311023-0323233031102210"></a>

<a id="canonical-3133300302112011-0121100112033010-2011023312100313-2212302312330022-2030131133202210-0332210103022210-1103130203232122-0000303200331333"></a>

## max_version property — custom_security / 313200221313 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1013030331321021-2120321112101012-3022010120221302-0111132203322312-2311132031001112-3313111301120021-2033233033101010-2300022301323211"></a>

<a id="canonical-2111300123030220-0213002003333123-2032221221222221-2310310001130002-3312030101112121-2132033210221210-0133002300103321-0213233230301132"></a>

## min_version property — custom_security / 313200221313 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000131001031120-0133120111112232-1222213333211210-1021013330211123-3110233211100213-1210110302013112-2313133222323123-2222130221303312"></a>

## Next pages — custom_security / 313200221313 / 7

- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0231221023332112-0013111300113021-1321302003220333-0312313301310013-3102001013231122-3231021230020303-2102021300103010-1121020000012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032122333123111-3220103302003331-1001331321033121-1133320021313333-1301221213030221-2012012323121111-3022301333300230-3322212032110033"></a>

## tls_tcp.tls_cert_params.tls_config.default_security — default_security / 111222201032 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- tls_tcp.tls_cert_params.tls_config.default_security

<a id="canonical-3000210203333221-0332202020313002-0322130023111302-2130322221231110-2231122122020131-3232223200011221-1130330030223203-2213333332001121"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0231132301211012-0012112213032223-2131231012121321-2032021230320023-2312010111232223-3021320302022201-1231320202203101-0302021000221232"></a>

## Direct properties — default_security / 111222201032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032131120103200-1303100311130122-1031100121132100-3323332111332030-0013333310232200-0332031230132033-1001032200120131-2201221322212132"></a>

## Next pages — default_security / 111222201032 / 4

- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3013211012113110-1100321110200123-3223300302021332-1222102203330102-3010300131111231-1000132312212100-1220213120001221-2222031113301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323012001231031-3300003033320003-1203323022011201-3233330231332221-3212202311122113-2102030022223200-1223220011230123-2111220301212032"></a>

## tls_tcp.tls_cert_params.tls_config.low_security — low_security / 031322020000 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- tls_tcp.tls_cert_params.tls_config.low_security

<a id="canonical-0030013133102302-3111331100103230-0212131123223001-2130232010030120-1320021111111002-3020012013103223-0313320310121232-2011202330222023"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0110202130032030-1300223233122013-3333100310010302-0210123301120103-1130312010011201-2231031230320201-3302100013232122-1301030121132023"></a>

## Direct properties — low_security / 031322020000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023330200003230-3230233321212323-1200110312233211-0300020230012002-1122011220310302-0121022131011001-1030030030221330-2312010010030333"></a>

## Next pages — low_security / 031322020000 / 4

- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2212002321101031-2100301010312210-0132001322000130-2120332323120010-2301113321223232-2223310031112030-1111100220031111-2121312120221200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110120332000230-3113022333313012-3103221030023222-1102031011200030-0231301331212113-1003231232301223-3100323113123133-0100210022232330"></a>

## tls_tcp.tls_cert_params.tls_config.medium_security — medium_security / 331302033213 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- tls_tcp.tls_cert_params.tls_config.medium_security

<a id="canonical-1012320000010113-2333030302211220-2313132232130223-2130011220132223-2031110103133103-2201032313000012-1122201103311331-3221213012232132"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1131331103213102-2232113012001211-0222220311012212-0032210123002000-0102301320103233-2313012230312223-2331032022113333-3212020111303001"></a>

## Direct properties — medium_security / 331302033213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311031123113331-1223330313233122-3231033030302230-1120303212302233-0011323002330300-3010331300113011-3131032022133200-0000201312221033"></a>

## Next pages — medium_security / 331302033213 / 4

- [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100302312002001-0302323333102000-2320110211310312-1221302023121111-0302110010301300-2030123103010303-0333032222222220-2022021001010132)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330012313003130-0323300023003131-0102223022220002-3100323122321000-3130203231023032-3321232312332300-0201132113033123-2213032312202022"></a>

## tls_tcp.tls_cert_params.use_mtls — use_mtls / 200310020022 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- tls_tcp.tls_cert_params.use_mtls

<a id="canonical-3012321301330313-1130000121202210-3132002213121330-2101110133000102-2221110132220130-0130220212232021-2133131310131330-1221300011113332"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-0201311320211333-0320313010011303-1003212122332300-1231323330320200-2200320221131110-2111220212330010-3202313303323133-2333130010220013"></a>

## Direct properties — use_mtls / 200310020022 / 3

<a id="canonical-3203000003301102-3110322001303011-0132221021312331-3013220331211213-1202220320222030-1010111030021230-1110231202231202-0100301210001233"></a>

<a id="canonical-2010323113131011-0332302103102302-3320312130133301-3130011121210131-2222220000012331-3112010120120331-3311030210333110-3321201310003211"></a>

## client_certificate_optional property — use_mtls / 200310020022 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2033220303231300-3223100331113101-3120220103323101-0203331202123203-0210133213200223-3102003112030231-1122320300211022-1301202022323010): complete subsection reference.

- [no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0313210121333221-3230321300020031-3302130032010203-2020100103212322-1133031103031310-0323223113032022-3200333330223322-2030313311321323): complete subsection reference.

- [trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1312123122202323-3133021133013122-2323012131313001-2230121233132130-2000110130212221-0230231002002133-2300301312012111-0010201102213322): complete subsection reference.

<a id="canonical-1232023003032313-1211121000333331-1101023310312120-2222112113233132-1030000122122211-0002200212103033-3103001212020322-0330231220022331"></a>

<a id="canonical-3213200000001231-1112030313021301-3222311203101011-1103111120222002-3123221220300212-3222013100311013-0302332020333100-0013011200011202"></a>

## trusted_ca_url property — use_mtls / 200310020022 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1232001321113131-3200223213223231-0031031230030102-0123303211211132-1213110111320003-1210222123313201-3233032203112130-0231311303221300): complete subsection reference.

- [xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3232023203013122-3310021000322212-3310003321033000-2002221122201102-0000102020010230-1331202232313223-0012110121010012-1120301201001210): complete subsection reference.

<a id="canonical-0033322231103000-3003113200201011-1021320312001231-2103331221211310-0312011111230012-0113231213332002-0113201203033010-0130023130122301"></a>

## Next pages — use_mtls / 200310020022 / 6

- [tls_tcp.tls_cert_params.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2033220303231300-3223100331113101-3120220103323101-0203331202123203-0210133213200223-3102003112030231-1122320300211022-1301202022323010)
- [tls_tcp.tls_cert_params.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0313210121333221-3230321300020031-3302130032010203-2020100103212322-1133031103031310-0323223113032022-3200333330223322-2030313311321323)
- [tls_tcp.tls_cert_params.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1312123122202323-3133021133013122-2323012131313001-2230121233132130-2000110130212221-0230231002002133-2300301312012111-0010201102213322)
- [tls_tcp.tls_cert_params.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1232001321113131-3200223213223231-0031031230030102-0123303211211132-1213110111320003-1210222123313201-3233032203112130-0231311303221300)
- [tls_tcp.tls_cert_params.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3232023203013122-3310021000322212-3310003321033000-2002221122201102-0000102020010230-1331202232313223-0012110121010012-1120301201001210)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2033220303231300-3223100331113101-3120220103323101-0203331202123203-0210133213200223-3102003112030231-1122320300211022-1301202022323010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333200323012121-0121211022112232-2232302311120111-3322302230131001-1221010022010223-0022210021321312-2213123300031232-1223222010222000"></a>

## tls_tcp.tls_cert_params.use_mtls.crl — crl / 221123131310 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.crl

<a id="canonical-1021021310222211-1220302113221211-0213213022332030-2020103012212202-0132311010011202-3003023002121221-3032010100221101-3230133203310033"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2033000221200111-2102220113221211-1110303133211202-0300222312103212-1032102313231302-2320101033103021-3133013003220100-0331013200320013"></a>

## Direct properties — crl / 221123131310 / 3

<a id="canonical-0013221000300021-1310113131223003-0112323010100101-2200312132322003-2033331121101121-0220013131002030-1101310130133101-0110011112320012"></a>

<a id="canonical-3323303211022032-1101230333030033-1310302103321332-3111301023213211-3303333011103020-0110102322211101-3201332130121021-1103321312223321"></a>

## name property — crl / 221123131310 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0113021201200122-2133120112123120-2111223032133301-3231032221031333-1102132313013120-1013300102301030-1233013331000322-2031121002111330"></a>

<a id="canonical-0113110020102003-1002312111212211-1001101213031002-1203323123302333-0231101021332320-0100200113033013-2102100220200311-1003010111102302"></a>

## namespace property — crl / 221123131310 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3301003122110202-2120301021221302-2232110131313300-2223321310230010-3113330233022230-3021210113132210-2123000333232012-3022031102230113"></a>

<a id="canonical-2002301103030111-1030110200333131-3131123101200111-2033213013103213-1032011032211200-0230132101032220-3332030322130100-2032331333123122"></a>

## tenant property — crl / 221123131310 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1101223223211011-1220102121112211-2111132132312222-2033333330233210-1003120121221222-0200210311322111-1303331121301001-0112100111203312"></a>

## Next pages — crl / 221123131310 / 7

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0313210121333221-3230321300020031-3302130032010203-2020100103212322-1133031103031310-0323223113032022-3200333330223322-2030313311321323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111111120133221-0011320333230022-1133030020101023-3021232301130122-2202333113310102-3121222212122003-3131231123311111-0022223031012132"></a>

## tls_tcp.tls_cert_params.use_mtls.no_crl — no_crl / 220030322021 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.no_crl

<a id="canonical-1313221332320211-1011200022102330-0322213303201222-0023330132023302-0212301222303112-3101213100332003-1001222201313113-2312020232003130"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0031212030332231-3320101011010303-1123100121213003-0201302312031310-3022123031333122-3322101012111320-0132233303320022-2212011113213030"></a>

## Direct properties — no_crl / 220030322021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330031113003202-2131123211231021-3130011222312201-1120233023122032-1102200130123013-0230130300222203-2203322000232030-3222122011303130"></a>

## Next pages — no_crl / 220030322021 / 4

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1312123122202323-3133021133013122-2323012131313001-2230121233132130-2000110130212221-0230231002002133-2300301312012111-0010201102213322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213333213110030-3010333022213313-3132203203010002-0001230023033112-2310102123033230-0002101332122012-1200103331330000-2330031221010310"></a>

## tls_tcp.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 010130330132 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-3100312211000133-0002100201332030-0312320012202300-1320103211122023-3122332010003131-2332321030001222-1122133120330133-2313011031120223"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0110322013223103-3300321231232303-3303023323010031-3211102201133102-0020232100332331-2113330210003003-2210002301111222-2013010232322033"></a>

## Direct properties — trusted_ca / 010130330132 / 3

<a id="canonical-3031030211121023-0330302030001022-3133223320233311-0122020313300322-2031210230323313-2211113133102110-0202223023133120-1321330133012103"></a>

<a id="canonical-3201103112303302-0213312103312021-2312130113200022-0030211111110300-0003030020033102-1031332020230003-2202000030001022-3100103123101101"></a>

## name property — trusted_ca / 010130330132 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2212121103233230-3211133303132030-3300213010321023-2021211120120203-2111120220001121-2030320203320010-1321330002233010-1023011322120130"></a>

<a id="canonical-2013223210110130-2210301110200203-1110313213233230-1120330231123203-2220321131001310-3001311003131213-2332303332330302-1020212102000213"></a>

## namespace property — trusted_ca / 010130330132 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2210122211323312-3012020130321013-0233130102203210-2030013133202311-2321230221121310-1321012003132212-2210132002133330-1300121323132302"></a>

<a id="canonical-1201102321132233-0233223002020322-3321002320133300-1002010213113101-3023200203331213-3213233032301102-1212010033212012-3202212320213300"></a>

## tenant property — trusted_ca / 010130330132 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3312203102322313-2123012102230330-1233301302330100-0133322111030330-2222320131102300-2132233210010112-3312212230230333-1320013220020231"></a>

## Next pages — trusted_ca / 010130330132 / 7

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1232001321113131-3200223213223231-0031031230030102-0123303211211132-1213110111320003-1210222123313201-3233032203112130-0231311303221300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001302012300332-1022030002011132-1101330302303120-0213210302020112-2313101232232300-2011002330301320-1033323100103231-1303033113030132"></a>

## tls_tcp.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 200022313222 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2012330212030210-1130211020301313-1030130013333121-0003003012210021-0023313032110211-3112133000211232-1030103033122100-2031121123003333"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1331223122213301-1303101231213220-3103213212012133-3313333220203122-1310232020221313-2133112113002331-3113021320011311-2110232102222110"></a>

## Direct properties — xfcc_disabled / 200022313222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321203220200213-2302203330130230-1121032001002300-1111310333012123-0211332032011021-2012031211220223-0132113132130031-3202223201221011"></a>

## Next pages — xfcc_disabled / 200022313222 / 4

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3232023203013122-3310021000322212-3310003321033000-2002221122201102-0000102020010230-1331202232313223-0012110121010012-1120301201001210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232233303011313-0322123132010213-0231302130203300-3131020203330203-0332010101001123-3221313310212301-0030211332303320-1100300331130212"></a>

## tls_tcp.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 023110231120 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1321302220102301-0123131312032311-3321222233330210-1031230021103103-1220002123300131-2013321333311331-2003220013332131-3022002201133123)
- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- tls_tcp.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-3113333131301233-1012113033202310-1200303321122130-0311033223320133-0300131110102203-0133233000013023-0112032032022021-3132023202202111"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-1222012131212331-0003112322233333-1110311220302312-3203002123202331-1021320313230320-0013013120213230-1332102100103013-2100321331011000"></a>

## Direct properties — xfcc_options / 023110231120 / 3

<a id="canonical-0223223130003202-3320131021110302-2323220000220122-1013112221003111-1100223112320022-1322233300031301-1132133221302111-3002032130322232"></a>

<a id="canonical-2130033222230130-0031232133201330-3132330030301030-0223121112231122-1313132233221333-3131021231212003-2113312013222030-1313220120303021"></a>

## xfcc_header_elements property — xfcc_options / 023110231120 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-1211321323211313-2310122112032313-1002133302132022-3203220123310101-2303322022212100-2321111223013202-0330112203023311-1121333231131213"></a>

## Next pages — xfcc_options / 023110231120 / 5

- [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330000323002301-3323302021320211-0333111311322213-1021311322130020-2131102031121120-3303220023112232-0000233302231300-1203021011230232)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011012320331123-2211132222210220-2320200331030011-3222323212112303-0200231220320133-3201032220311033-2313010020323132-3102212311230011"></a>

## tls_tcp.tls_parameters — tls_parameters / 010202333310 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- tls_tcp.tls_parameters

<a id="canonical-2121021311101012-2230010312001123-2312121013311221-3322320113212022-1023322132223233-0233320112310313-2000233212123201-3101323233030001"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-2230312133210222-3233132111212211-2110311112230230-0100220000301303-2111022033300123-3221323102003101-2023312322231230-1230131101121303"></a>

## Direct properties — tls_parameters / 010202333310 / 3

- [no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0233013330312022-1322221102311201-1202300132103012-0133230011223300-2131202330320131-1210323030103011-2000210010321102-1223113020012032): complete subsection reference.

- [tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333): complete subsection reference.

<a id="canonical-1332133320010033-1232010221213330-3010333020313312-2223331013133133-2231211303000310-2333012221023213-2122203010112223-3203001223023300"></a>

## Next pages — tls_parameters / 010202333310 / 4

- [tls_tcp.tls_parameters.no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0233013330312022-1322221102311201-1202300132103012-0133230011223300-2131202330320131-1210323030103011-2000210010321102-1223113020012032)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0233013330312022-1322221102311201-1202300132103012-0133230011223300-2131202330320131-1210323030103011-2000210010321102-1223113020012032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021023331023122-0113310200133211-0010130303312121-2132131312000102-2110323223122003-2110310111313011-3302100032232010-2332021102132012"></a>

## tls_tcp.tls_parameters.no_mtls — no_mtls / 021302102103 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- tls_tcp.tls_parameters.no_mtls

<a id="canonical-3001133123312002-1000320221230212-0222121000310303-1200233122223222-2230332032132333-3313210001031231-1313031322200131-1021111110011000"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0012122221131231-2221332120130210-0032122321023330-2031002311103110-0333102131201102-1203333211323121-2021223132302012-0221220131102213"></a>

## Direct properties — no_mtls / 021302102103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220311011020331-0230123201030121-2102332313213311-0212302221100311-3222323000031123-2311000013321311-1201100113010220-2120131102231012"></a>

## Next pages — no_mtls / 021302102103 / 4

- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222312330033003-3131213030102020-2132232122212112-3121313222200001-3103233312323110-0332013132133100-1330201220112322-0331122331133131"></a>

## tls_tcp.tls_parameters.tls_certificates — tls_certificates / 303030131203 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- tls_tcp.tls_parameters.tls_certificates

<a id="canonical-1012102230213233-3212321103023301-1031112200203333-0003130113203002-0113233211132130-2100101021100331-2303203302000310-2130011002231231"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3003033121322100-0122313312221331-2103202110020033-1132113203302332-2130320021130233-3011033223332131-1123221303322331-1132302200312121"></a>

## Direct properties — tls_certificates / 303030131203 / 3

<a id="canonical-1100213020121312-1102013313211022-2021013123233101-1321231012321020-0320132023320322-1303211332030232-2331203210331133-0332202232232121"></a>

<a id="canonical-1033321121310220-1033120130213331-3132310211202101-3003202221120203-0203131203013311-0230223301301103-1202201200332012-2122122203031310"></a>

## certificate_url property — tls_certificates / 303030131203 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1312320021030200-1023033111010123-2231303213022200-2333113221031111-3331010212013101-1221030110303211-1311330102003211-1303112122010003): complete subsection reference.

<a id="canonical-2230333001233120-2021110211021230-1132332320103130-1112110302310001-2322212112121030-3310320120310020-1331013233132303-2210320312222033"></a>

<a id="canonical-0012331002132223-0002103101100113-2321122310033313-1031320311311022-3120222001110231-3230232322012020-1300220210030213-0201301013210020"></a>

## description_spec property — tls_certificates / 303030131203 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2211310213312032-2023101020123311-3223220131132300-1013031022211013-2223221012220020-1231010030320333-3300313033101131-1300230323330023): complete subsection reference.

- [private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102): complete subsection reference.

- [use_system_defaults](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3013202212033200-2301011303032213-1113001000221013-2321021112001011-2333001300303322-1121311110012323-1330012011103030-0230031010332022): complete subsection reference.

<a id="canonical-2223323121210021-1103011332021333-1013032123202131-0123311233211031-2223022323332001-0330210300200203-2212232031303300-1212101202301312"></a>

## Next pages — tls_certificates / 303030131203 / 6

- [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1312320021030200-1023033111010123-2231303213022200-2333113221031111-3331010212013101-1221030110303211-1311330102003211-1303112122010003)
- [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2211310213312032-2023101020123311-3223220131132300-1013031022211013-2223221012220020-1231010030320333-3300313033101131-1300230323330023)
- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102)
- [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3013202212033200-2301011303032213-1113001000221013-2321021112001011-2333001300303322-1121311110012323-1330012011103030-0230031010332022)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1312320021030200-1023033111010123-2231303213022200-2333113221031111-3331010212013101-1221030110303211-1311330102003211-1303112122010003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333223211123312-0001232000301323-2232001213333113-2102313300201003-1031312032133031-1002111220323132-1331101200130022-2233311302132300"></a>

## tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 012002002222 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-3333312302220231-2222020031203330-1203200113121021-2201210323030201-3033223233113213-2302123310131210-1213103311101220-1000102233110122"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-0020130132133223-2212131312121230-1321200033332223-1030201033130213-1003131101000001-1132002223321100-2130232323233023-0323202032000231"></a>

## Direct properties — custom_hash_algorithms / 012002002222 / 3

<a id="canonical-0201303311112033-2123130100113320-3331110213002110-1030033210221302-0331221313203021-1322222013121220-3312222013132233-1022000102330003"></a>

<a id="canonical-1120103203023111-1003203220000032-2003231031131131-0210030330101112-0322323313333321-0330232310222330-2013233112220201-2111303223211020"></a>

## hash_algorithms property — custom_hash_algorithms / 012002002222 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3321213313310122-1002303222221003-2213310322103030-0123332333312002-0210300101031220-0020123233210301-0230231030101210-1120120000312013"></a>

## Next pages — custom_hash_algorithms / 012002002222 / 5

- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2211310213312032-2023101020123311-3223220131132300-1013031022211013-2223221012220020-1231010030320333-3300313033101131-1300230323330023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111200231000023-3210220321022032-0231332131030233-2022311120032112-0311013223101201-1000113213233122-2121122320210212-1212300112312101"></a>

## tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 311233331001 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-1011133323230033-0011110010130221-1323003103012012-0103311200301030-1130220220222033-2103111221122030-1320001303310010-2021201132022322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1310112313223322-3231300222131003-2322020323332100-2212102101320233-2111111220021001-1201131011031233-0103013331021321-1300232220333211"></a>

## Direct properties — disable_ocsp_stapling / 311233331001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122133303112212-1311121220023323-1220031133103323-0222322333202301-2131211232102122-0100301111211323-3222331011200103-3320112032231031"></a>

## Next pages — disable_ocsp_stapling / 311233331001 / 4

- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311121003122132-2330202003301322-0332021231020200-3013003200131133-3203010212330231-3231300312300221-2023211220132032-2032001023122333"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key — private_key / 221102010111 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- tls_tcp.tls_parameters.tls_certificates.private_key

<a id="canonical-1101331003131213-1231312032211322-2233011011231000-1002211131023110-1022021311003332-0223300202331220-3102031301321103-2220022111213132"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-1211000221213113-3122320130230302-1031303031313332-3020312213130310-1132131301123123-0130111201020223-2333221131020112-1100333002313120"></a>

## Direct properties — private_key / 221102010111 / 3

- [blindfold_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1313303320012120-1133213001033103-0320333003301302-0200033133331101-3312030010123122-1001133223300000-2213213322101003-1000112231200000): complete subsection reference.

- [clear_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0210110131100012-3202013230230003-2233100231111021-3233201333010332-1312220310320112-1003101221321103-0102213130030101-0032322120313211): complete subsection reference.

<a id="canonical-1221030122231300-0332023231031113-0201111230222302-1321111100330300-2332201321312021-3112112221022031-1332002331111311-1100002111010101"></a>

## Next pages — private_key / 221102010111 / 4

- [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1313303320012120-1133213001033103-0320333003301302-0200033133331101-3312030010123122-1001133223300000-2213213322101003-1000112231200000)
- [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0210110131100012-3202013230230003-2233100231111021-3233201333010332-1312220310320112-1003101221321103-0102213130030101-0032322120313211)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1313303320012120-1133213001033103-0320333003301302-0200033133331101-3312030010123122-1001133223300000-2213213322101003-1000112231200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222020203212300-0202102322021222-2000122230330332-3102233312031200-2220322121131333-2202332121112331-0113032011103231-0102211032023123"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 103223101220 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102)
- tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3320320232322331-0001011010110300-2033233300111122-0013003310303200-2323210133200202-3030100013323002-0021300202230233-1122010002233112"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-2231211120301231-3233333210210223-0130233100100032-2033300023103022-3233110120333232-3031222301030300-2002230030032322-1222332201001311"></a>

## Direct properties — blindfold_secret_info / 103223101220 / 3

<a id="canonical-3033010312023320-3132320200102133-2303231033212110-0012020102201103-0120213302220303-2211210332001202-1301232302221232-3221221323301100"></a>

<a id="canonical-3303322310010132-2330333121313010-1122010231012033-2113133200213132-3212212030010102-0211333102033321-2101232310012011-2100113131321233"></a>

## decryption_provider property — blindfold_secret_info / 103223101220 / 4

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0133330130022222-0021302211130321-3023230232220112-0110000232322331-2201321323012013-3103103132211023-1021332230220130-0310013333331303"></a>

<a id="canonical-3033131122102223-3130101032203100-1311002312100033-1001310230032023-2220303021221010-1003321332330213-0333223330021202-0221112303013210"></a>

## location property — blindfold_secret_info / 103223101220 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2023322332202320-2213301030232110-0302133102330320-3230222011123313-0120333221212013-2003103222201320-1320131031202012-1031323022221231"></a>

<a id="canonical-0012220033331310-2233201230311132-3130221233020132-3013130313001131-2120230313320103-3121221300013312-2210332113311002-1002023210321312"></a>

## store_provider property — blindfold_secret_info / 103223101220 / 6

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3022323021303113-0300232300032330-0102122210122231-0100323231133331-1021021212220311-1303311313313123-0011033230101330-2232101201231321"></a>

## Next pages — blindfold_secret_info / 103223101220 / 7

- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0210110131100012-3202013230230003-2233100231111021-3233201333010332-1312220310320112-1003101221321103-0102213130030101-0032322120313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010002003302121-1220221213321123-0202102030201302-3120013213120022-3313132223023131-1320003032233023-0020122322230130-2310003310313001"></a>

## tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 121210122300 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102)
- tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-1312210322322010-2311302023001232-0032233233302100-1121300131032223-0232313303000000-0132231331200033-2310111021202313-3011303222210213"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-0002310232222320-2020321033312131-0123100033033300-1310121331013001-3230032320313202-3122110323101213-2111303213001022-0010110013321103"></a>

## Direct properties — clear_secret_info / 121210122300 / 3

<a id="canonical-3003112000123322-1131112122120313-1113332232201322-2032211033000333-3313120132010033-3130233333012003-1113033130331320-1011201223003101"></a>

<a id="canonical-3130001300032010-2002023002013201-1121233301301021-0031010103120310-3100031230233113-2223233021211323-0211212322313312-0232102210320010"></a>

## provider_ref property — clear_secret_info / 121210122300 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1023030133133323-2302200000330013-3220112010000021-1221303222033010-2110120301010220-0131220221211213-0313101120011321-0003322102212211"></a>

<a id="canonical-0132303002203303-3201303212330110-2100223001133013-3101321133020322-1110110023002111-2003312002201331-3032130330200120-1220302312311213"></a>

## URL property — clear_secret_info / 121210122300 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0332211333222211-0320302001231221-2022223321002202-3202012032332122-0002032030032200-1003313320321303-2312012000320311-2022223312320300"></a>

## Next pages — clear_secret_info / 121210122300 / 6

- [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0110223113120001-3133220231303130-1303100213120130-0021230232101221-3110100100031300-1320002101200103-1122012331121220-0130220112211102)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3013202212033200-2301011303032213-1113001000221013-2321021112001011-2333001300303322-1121311110012323-1330012011103030-0230031010332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103331012330331-1001232112213310-0122033002120111-2301000333220032-3333311203120310-3110200030310333-0031232202013111-3211103010021021"></a>

## tls_tcp.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 211011033332 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- tls_tcp.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-0133232121100203-3112131312222332-1210211201023203-0213031301113222-1100110303101231-1002321023220200-2130320023203200-2102003202313001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2111200032330010-2022010201111211-1130023312123133-1211113022221303-3232002111010121-1111003213031313-0322220121221302-1300121201103033"></a>

## Direct properties — use_system_defaults / 211011033332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212332320323132-0210031311033322-1020031230130111-1130002320113132-2221003222220111-2113032232223113-2003202002032223-2102011112013012"></a>

## Next pages — use_system_defaults / 211011033332 / 4

- [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1013221020220222-1030002212300101-2323120001023021-1022203013021002-2310100211231130-2000212020333203-0233330031221232-1222100313012132)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330131231023313-2100231003001221-1310230332203001-3211131302002120-3000000013133010-1131222221132333-2320313213313321-1211222123033000"></a>

## tls_tcp.tls_parameters.tls_config — tls_config / 102012322123 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- tls_tcp.tls_parameters.tls_config

<a id="canonical-2221303331133203-1130203323222012-3330132321032001-2221230001200013-3211320212313313-1330022310200131-2012220011310103-1123313133101101"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-3023210131110323-1201101232223201-2312333220033201-2231311110332221-3202303011322011-1111131202321321-3231100022232303-3223102303220333"></a>

## Direct properties — tls_config / 102012322123 / 3

- [custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2313330220103123-2201220320310123-2011113302301211-1321222013332010-1231123323023011-3003130012232220-1002222233203333-1011000312201213): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1202210311313111-3213122012121030-0112132213012202-0212010323332023-0110123212333122-0000212302331210-3133120103320012-3022311123003121): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3320233103032110-0222202323122111-3120211200010322-0120203132021002-3031133233021303-2320223210003312-3030020302103031-2011202203323302): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1331020023012022-3032213132000232-0121020020323021-3101333010132310-2120022302203221-2112321202330321-3221300002133332-1313023302320030): complete subsection reference.

<a id="canonical-3020020122202132-0023201003101002-1000111103031020-2110122001032201-2021232110303020-0201110112322230-2002121310010123-1000203020320010"></a>

## Next pages — tls_config / 102012322123 / 4

- [tls_tcp.tls_parameters.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2313330220103123-2201220320310123-2011113302301211-1321222013332010-1231123323023011-3003130012232220-1002222233203333-1011000312201213)
- [tls_tcp.tls_parameters.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1202210311313111-3213122012121030-0112132213012202-0212010323332023-0110123212333122-0000212302331210-3133120103320012-3022311123003121)
- [tls_tcp.tls_parameters.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3320233103032110-0222202323122111-3120211200010322-0120203132021002-3031133233021303-2320223210003312-3030020302103031-2011202203323302)
- [tls_tcp.tls_parameters.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1331020023012022-3032213132000232-0121020020323021-3101333010132310-2120022302203221-2112321202330321-3221300002133332-1313023302320030)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2313330220103123-2201220320310123-2011113302301211-1321222013332010-1231123323023011-3003130012232220-1002222233203333-1011000312201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213320132210202-1223021321222203-2222310220001301-3211232232330330-3001231323133332-2300022021331332-2122013110021230-1312111031302330"></a>

## tls_tcp.tls_parameters.tls_config.custom_security — custom_security / 122333103101 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- tls_tcp.tls_parameters.tls_config.custom_security

<a id="canonical-3013101131313113-0201013023331222-3233331012300022-1301301333232231-3312223021100102-0312131001221132-2103123100122010-0102133012331132"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-1222123001232102-0123101001213331-1132130210120010-2112220121323123-3122012122001310-0103322232001032-0211201113002120-3231113232330212"></a>

## Direct properties — custom_security / 122333103101 / 3

<a id="canonical-3100112232003232-2300221211320311-2021232303000202-0300221100013232-0311030213023021-2013322320033113-3302320120122101-1310220133233300"></a>

<a id="canonical-1302110131101121-2222301310001313-0003332230323331-3113102213022003-3302313000231000-3222020100123212-3311011130011030-1023133310110213"></a>

## cipher_suites property — custom_security / 122333103101 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0300113000203223-2220033311213120-3310333103003013-0301032103312203-0011301303010303-2001211122103231-0111103021111010-2210022200322330"></a>

<a id="canonical-1123320132222023-2023111132003000-1120232131022032-1230313113000033-0112233333133000-0330123313202313-0002112020002331-1202321001213203"></a>

## max_version property — custom_security / 122333103101 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3331100231131101-2031232201123000-1032022131120020-2331330012113233-2331102112333131-2202232211011330-1202011102333313-1110203132310202"></a>

<a id="canonical-0012033032123332-2212321102301032-3333111231000100-0312021131002103-1211212010123111-1130031123311222-3112012030331302-2221223301001331"></a>

## min_version property — custom_security / 122333103101 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0001123011121212-1311320000202201-0122022031000021-0322222112131112-0233321010132030-1000000302212203-2323213030202111-3102021020213231"></a>

## Next pages — custom_security / 122333103101 / 7

- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1202210311313111-3213122012121030-0112132213012202-0212010323332023-0110123212333122-0000212302331210-3133120103320012-3022311123003121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310213021313023-2333210113100300-2310023103323231-0120131100221322-0102121030101020-1222323211232003-2112003133100221-0223120021320332"></a>

## tls_tcp.tls_parameters.tls_config.default_security — default_security / 131210232103 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- tls_tcp.tls_parameters.tls_config.default_security

<a id="canonical-1332022222321223-2110230113333020-1021123232000223-3223023231122321-0301011020323100-3232031023123112-2223012332033230-2012203201302001"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2233111032002132-1123001331212133-0302330313001322-2323012300222213-0022202211113320-1021201023300320-1011103123002310-2122310231101321"></a>

## Direct properties — default_security / 131210232103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230233231313111-1211222322223031-1021210001022222-3220323320303321-0232330233230113-0113223001222131-1323112212022212-2033011023031030"></a>

## Next pages — default_security / 131210232103 / 4

- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3320233103032110-0222202323122111-3120211200010322-0120203132021002-3031133233021303-2320223210003312-3030020302103031-2011202203323302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303330220110000-0100103331321302-2332302020120230-0321211230223313-2013133003132212-1110212203022210-3203221230032023-2210033130302000"></a>

## tls_tcp.tls_parameters.tls_config.low_security — low_security / 333310201231 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- tls_tcp.tls_parameters.tls_config.low_security

<a id="canonical-1233020333030001-0023231023111002-2320211210113132-0200321030010330-0222301121001232-2031100331121301-0100303211321012-3311212012121130"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1301002103201102-3021010100001103-3230230003002132-2033232030012112-3233132110231310-0203032122202122-2211312312030132-3102311201303133"></a>

## Direct properties — low_security / 333310201231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322010021330032-0233013223232220-0132332013201012-1310131021010332-0213213122023211-1332330112302110-3213030223331230-2101210023223113"></a>

## Next pages — low_security / 333310201231 / 4

- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1331020023012022-3032213132000232-0121020020323021-3101333010132310-2120022302203221-2112321202330321-3221300002133332-1313023302320030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123302313300112-1200323003031323-2231223102221321-3111220022023311-0333031020333231-2023221003030033-3022033020033120-2011113232030323"></a>

## tls_tcp.tls_parameters.tls_config.medium_security — medium_security / 133232110321 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- tls_tcp.tls_parameters.tls_config.medium_security

<a id="canonical-3001103131311223-3010020313031202-3200003032201111-0002112323322103-3123310233032000-3131111123113132-3102113330310000-3231122221311101"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-2302210322331203-1230130321120302-1232110331001002-3101220023033130-3021131321232200-1030200300133102-0012333112111232-0032131211012230"></a>

## Direct properties — medium_security / 133232110321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303211213101310-2320000032002102-1000100130322310-2010201203303233-0312322031012020-0223222003303113-2332231311330102-1121020133030110"></a>

## Next pages — medium_security / 133232110321 / 4

- [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113122011111202-3302111302113203-3031002313111223-1323002203100210-3200110320233330-2211220331301211-3313123232000210-1021322223230011)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303311113012301-0232130033301002-0013021321000203-2221030213002320-3003310201331100-3031121300213013-3131233330120113-1010020331121330"></a>

## tls_tcp.tls_parameters.use_mtls — use_mtls / 130221230133 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- tls_tcp.tls_parameters.use_mtls

<a id="canonical-3223222113230022-2123203033302010-3201002232021200-2112303222332011-1013320230313022-2100203313201023-0320013033033023-2203133121123123"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-1011213212223123-0102212011200300-2211021012122212-1101213331102123-2303003320001132-1303132211121112-3023102203230112-1300133222100223"></a>

## Direct properties — use_mtls / 130221230133 / 3

<a id="canonical-0310032321121022-1131300132301122-0101302232113203-1330002133211013-1022010011303221-1310002003321301-2013330120233011-2331230323113310"></a>

<a id="canonical-2211100121120233-2030112202302001-0232222013103311-3333230131001320-3300013200332010-3011112011032122-1133301311101033-0321021031121120"></a>

## client_certificate_optional property — use_mtls / 130221230133 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3113103210011021-0023200113021303-0030130323322102-0132102302111102-3320331113032113-1222212231001332-2320122112023030-0110011013320110): complete subsection reference.

- [no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0231000011022020-0131323132201333-0233202001110131-2112133122102301-2023333032022121-3100321310331130-1223300310301230-3133100232113130): complete subsection reference.

- [trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3203103003231323-0010301020132033-1312333201232130-3313333321111313-3203333013300103-2122333003332312-1112012110001020-2223201310331302): complete subsection reference.

<a id="canonical-3011301223002313-1302132132023313-0330120020322202-3032302202111112-0010320111013303-3200332330210301-2302112231221303-2112120130322120"></a>

<a id="canonical-1030110000102130-0213002000010002-2120303331113311-3310132101330000-3220333211323030-1121231232330121-3221321222012220-2010331233202233"></a>

## trusted_ca_url property — use_mtls / 130221230133 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1310312103233030-2113330332320313-0130130222220301-2103021121021111-0123113011232102-1130221201123030-3232312231303131-2332100210012133): complete subsection reference.

- [xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0102203131003320-1313021231132023-3111101010120312-0213011120320200-0032102200003313-0133301101301332-1012333131320011-0012003103322233): complete subsection reference.

<a id="canonical-1101303023220300-0213031120030002-3011310103302000-0200323213103120-0113310302101301-1203130321122031-1322320103020222-2002302002110233"></a>

## Next pages — use_mtls / 130221230133 / 6

- [tls_tcp.tls_parameters.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3113103210011021-0023200113021303-0030130323322102-0132102302111102-3320331113032113-1222212231001332-2320122112023030-0110011013320110)
- [tls_tcp.tls_parameters.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0231000011022020-0131323132201333-0233202001110131-2112133122102301-2023333032022121-3100321310331130-1223300310301230-3133100232113130)
- [tls_tcp.tls_parameters.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3203103003231323-0010301020132033-1312333201232130-3313333321111313-3203333013300103-2122333003332312-1112012110001020-2223201310331302)
- [tls_tcp.tls_parameters.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1310312103233030-2113330332320313-0130130222220301-2103021121021111-0123113011232102-1130221201123030-3232312231303131-2332100210012133)
- [tls_tcp.tls_parameters.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0102203131003320-1313021231132023-3111101010120312-0213011120320200-0032102200003313-0133301101301332-1012333131320011-0012003103322233)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3113103210011021-0023200113021303-0030130323322102-0132102302111102-3320331113032113-1222212231001332-2320122112023030-0110011013320110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033013301212213-2331223123032203-1320020123112203-0023233023233101-1333020111021200-2322133003200200-3221200203021311-1031012022123002"></a>

## tls_tcp.tls_parameters.use_mtls.crl — crl / 020202122020 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.crl

<a id="canonical-0100200111001303-3031030302320223-0210132001100020-0232033032122220-2023302130333033-3322212220301321-2230201112022322-1103221030212303"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2103232102200100-0221123302212213-2133000300030003-2200201213130322-3331313111223023-3322311132133313-2111200023011233-0012032121223221"></a>

## Direct properties — crl / 020202122020 / 3

<a id="canonical-1112302322302231-0333230320023333-1313222210010222-3020100210320032-0212133121233013-1202111100012021-2121001322020311-0033131210003102"></a>

<a id="canonical-0230120330013222-1002300320201312-0300002110111323-0221120013332002-0033302003102311-0310203002030012-2302033013301302-0331232011333011"></a>

## name property — crl / 020202122020 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0101102323230331-1033111200123212-0313133311121003-3302110203330022-2312030020033123-2013312213303022-1332232331231313-0231323130312312"></a>

<a id="canonical-0003111202313212-3221321020333000-1211333132110103-0232210112122230-0302100002101222-2112010212102112-3232203113101200-1310113210022130"></a>

## namespace property — crl / 020202122020 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3121333021032331-0112221131203000-2230102322330110-1212202010203110-3312012231213313-2033322210322331-3323302330310322-1301003203021132"></a>

<a id="canonical-1310323213321103-3222331303213110-0312132102233310-2022313230201222-0013132321100003-0133202030133302-2232113331012323-1213010102101313"></a>

## tenant property — crl / 020202122020 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2313303220122023-0323113112321130-3113122333211012-1123011033031002-1113131332333212-2121012210310111-0211021020322331-3120133213233102"></a>

## Next pages — crl / 020202122020 / 7

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0231000011022020-0131323132201333-0233202001110131-2112133122102301-2023333032022121-3100321310331130-1223300310301230-3133100232113130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332110223122112-1031311000232222-2321303330121223-1030123000300021-0101101300310113-3101202031010230-3333300003202200-2320321022220301"></a>

## tls_tcp.tls_parameters.use_mtls.no_crl — no_crl / 123111302201 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.no_crl

<a id="canonical-0332112123301220-2310300121111102-2203101012332220-2313031111331131-0001313301002011-1101331312010313-2130222002312123-1233012313332223"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-1203301322130312-3230003201332032-1223031120002002-1101303103033200-1013001031103021-2331213011211132-0220312030120231-3032002110020012"></a>

## Direct properties — no_crl / 123111302201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201232331130132-0100331012223012-2131010211120310-1003202110012101-2202123102221112-2021133233010123-0023001213231100-1032101211122012"></a>

## Next pages — no_crl / 123111302201 / 4

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-3203103003231323-0010301020132033-1312333201232130-3313333321111313-3203333013300103-2122333003332312-1112012110001020-2223201310331302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131333012022210-1332031133013313-0010123010201002-0133111030310000-1302303102223211-0132011332133031-3003212210310213-3231211133132211"></a>

## tls_tcp.tls_parameters.use_mtls.trusted_ca — trusted_ca / 122123130032 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1310022002220230-3301211110111003-0120122110301201-2213301110010223-1231112203210021-3301302011233230-3001200111102203-2023313200231030"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2212013100210222-1033012320322332-2131323301302021-2111231232022121-0032213311033320-0331030321331122-2320332000130212-1222312200021233"></a>

## Direct properties — trusted_ca / 122123130032 / 3

<a id="canonical-3001302011113030-0232101312023210-1101333020331021-0000310222332101-1030013002322101-3333301122301213-1011203202223312-1101221023101233"></a>

<a id="canonical-1301113201322301-0110111122333201-2301003120221020-0002233011000303-2332003200012131-1131221133331200-1013031231133312-0332210133030121"></a>

## name property — trusted_ca / 122123130032 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3012232013130300-1211110120202010-1033320132320030-1221001200130331-2212033003210322-0333131030330203-3003113322003101-1301032022120211"></a>

<a id="canonical-0331113033210313-0332011231120231-1113223301123332-2012301021101103-1003311321122102-3110310133311233-1000330221221132-1331300030100213"></a>

## namespace property — trusted_ca / 122123130032 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3102121312210012-3230223000230330-0002200132121220-3303012000200202-2231302212012211-3202203102303302-1112112331221013-2332100132203313"></a>

<a id="canonical-0003222132320003-2302003223102300-0300102003100320-2331011332010233-0331302002313321-2323331133223203-2333222113333303-2200223230210012"></a>

## tenant property — trusted_ca / 122123130032 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3310022233001030-3011123113310020-0211123010213311-0010013211131020-1221201132130122-1311223202102032-0313213331121023-0000021022033012"></a>

## Next pages — trusted_ca / 122123130032 / 7

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-1310312103233030-2113330332320313-0130130222220301-2103021121021111-0123113011232102-1130221201123030-3232312231303131-2332100210012133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102023101103123-0312302231331103-1123222002230301-2220002220031010-2232230322022010-2121123223022213-2231323112320310-1022230133130110"></a>

## tls_tcp.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 010110003210 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-3120300001022020-3010321310302003-1230131212331113-2120233033020121-3202321332133232-0122331131222200-3313301220303133-2033330230300020"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-0012012020330032-3303312031131120-3031030323032111-0211200211332003-1130222303011022-1303110321030231-2122022133132132-1132021210100112"></a>

## Direct properties — xfcc_disabled / 010110003210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220311321020021-1033300310201222-2031213311120233-1030111311022113-0021113220011020-0021110131032031-1212323123120120-3022101322002033"></a>

## Next pages — xfcc_disabled / 010110003210 / 4

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0102203131003320-1313021231132023-3111101010120312-0213011120320200-0032102200003313-0133301101301332-1012333131320011-0012003103322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111312231323211-2032123133321333-1211222300200112-2320010003010310-1121022203321331-3112210112321103-2120110022313113-1003332223121003"></a>

## tls_tcp.tls_parameters.use_mtls.xfcc_options — xfcc_options / 010321310222 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301)
- [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1221312001300312-3200002331002010-1113200020313010-0000121300320330-1233223032110302-2222230031301332-2202120101001130-2330031333113120)
- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- tls_tcp.tls_parameters.use_mtls.xfcc_options

<a id="canonical-1020130210233233-3331300223301103-0110222032000010-1030001032212101-3020221332323101-1230201030332102-1112303223213101-2100020221011223"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-1301101311022333-0223000210120031-1000333132120302-3021010210322022-1002122031112301-0110220310012000-1002000112130221-2221213102111301"></a>

## Direct properties — xfcc_options / 010321310222 / 3

<a id="canonical-2203003223223132-0103012231300230-2103111213131010-0223322310201003-0300102321200222-1122132202030210-0230003011110302-1112111023231330"></a>

<a id="canonical-2213020132120020-1131123233123031-0200101301210221-1220100023322123-3211231233032230-3311220313202013-1212031111232102-3212231003032123"></a>

## xfcc_header_elements property — xfcc_options / 010321310222 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-3321032221101331-2232101311302212-3113120013130312-3230100012133022-1032120031230203-0011213232221222-2213030303311002-2112030333123022"></a>

## Next pages — xfcc_options / 010321310222 / 5

- [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3121102020301213-2031321122303113-0000213133001210-0210232230033300-2021332132320000-0222122132332101-0203010102112222-0202212331011333)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001221002101232-2110033203133201-1001101101021320-0332030212221112-0031202100200130-0321222010330123-0203020110223012-0313112331300222"></a>

## tls_tcp_auto_cert — tls_tcp_auto_cert / 303221330130 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- tls_tcp_auto_cert

<a id="canonical-3122010013013333-0010120031233203-2230121020320021-2303221102121132-3030013013032022-1321333312130213-1313113220130030-1002023233330000"></a>

Type: `"single"`. Computed.

Choice for selecting TLS over TCP proxy with automatic certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-1011123030300033-1231300022012210-1010001103101310-2002310302320000-0110200033002002-3213031220023322-3233110300120232-0220313321110023"></a>

## Direct properties — tls_tcp_auto_cert / 303221330130 / 3

- [no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0020031100213200-1201323322020133-1102331111203000-3232323102121231-3320311130130232-1110312330122333-0221033220013000-0203122030013300): complete subsection reference.

- [tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223): complete subsection reference.

- [use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122): complete subsection reference.

<a id="canonical-2312202223110032-0220323221301032-3320233003100320-0013130001103310-2333010312123103-2020331302130221-2030003233322021-0021130111003010"></a>

## Next pages — tls_tcp_auto_cert / 303221330130 / 4

- [tls_tcp_auto_cert.no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0020031100213200-1201323322020133-1102331111203000-3232323102121231-3320311130130232-1110312330122333-0221033220013000-0203122030013300)
- [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1001031123221213-1211323113223112-3110230023122012-0311333202102302-3102321013231233-2233121202122222-1303123201210011-1201202313331122)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-0020031100213200-1201323322020133-1102331111203000-3232323102121231-3320311130130232-1110312330122333-0221033220013000-0203122030013300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120120232322012-2301133220122301-0231331023133303-2221000322012320-0110130313200122-2010113001211323-3032333030023232-0101002011100120"></a>

## tls_tcp_auto_cert.no_mtls — no_mtls / 312210333332 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- tls_tcp_auto_cert.no_mtls

<a id="canonical-1312323322322321-1230033311202000-0021330230212001-2121300030011302-1310011212312321-2321022223213330-2022210331101330-2232332123301233"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-3010321033233101-0322202330232230-3003311321013013-2201302322031033-0220220013200023-3210212101303002-2213223301330332-2133333210001200"></a>

## Direct properties — no_mtls / 312210333332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300231021320033-1010103021210130-3312032130321322-0311213122332211-2010030000221323-0122002101212202-2113110010113330-0001030103132322"></a>

## Next pages — no_mtls / 312210333332 / 4

- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)

<a id="canonical-2030011130201210-0021230212213033-1322111102100230-3001311031201101-2332313133030232-3120203023113310-3010203232113230-3011130323002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110320230121221-3212210332303333-2110311133133323-0032301103032232-2301312130102321-3232012313200212-0023110123213121-2112311123102220"></a>

## tls_tcp_auto_cert.tls_config — tls_config / 230320200133 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221)
- tls_tcp_auto_cert.tls_config

<a id="canonical-2100223333301213-2002332220230023-1111133020100000-1303221120200120-2202022230303211-1102203132031311-1102301231123233-0210203300120303"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-2030310011232232-0112200310032003-3201101130100101-0333123230322302-0213003101220310-3311230332312010-0302131320002300-1323213203203131"></a>

## Direct properties — tls_config / 230320200133 / 3

- [custom_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3120310133012112-0130123232123221-1322002121223002-2320133321120133-3001330312233013-1201132210000310-2010200313230300-3333102030011100): complete subsection reference.

- [default_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1011222301233103-0101132330111200-3002231332032030-2203231131213312-1211210211020231-3200103321020100-1202011022023230-0032320221323302): complete subsection reference.

- [low_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3223122001220111-2210032122330221-0020001102233131-3130011002220131-3311122323002312-1112210212322202-2312211202121331-3013210100112313): complete subsection reference.

- [medium_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0110130132333313-3000221311101333-1312011230032310-0120312223201000-1101313331312023-2133033110100221-2320313021012030-1212212201231323): complete subsection reference.
