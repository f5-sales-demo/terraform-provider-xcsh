---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1100310113021011-2300121112223230-3322231112002312-1010012320233101-3131033110123112-0010231221033102-2012213121201332-2321131202330003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-0200002222113033-0202331111220101-1122122022313311-0012110333111330-0023300203321132-2013031310031333-0331302322210222-2030201112202301"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-1232300110032012-1132221313133331-3020221130222001-1202021230222110-1030010330221203-3012013321232030-1020030202022211-2022301031112332"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310111233113013-1012032030330032-2001031233320301-0120212223302220-0212231210110121-2112230133122213-0010200121002100-3121331123130332"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates`

- [blindfold](resources--workload--reference--group-010.md#canonical-1202110023222001-3100313121312113-2233103210010221-3332332021103000-2331000312013011-0212033213020210-0112201202301013-0010022021211032): complete subsection reference.

<a id="canonical-0310302211003333-0300220330123103-3301000231222130-1301000312332033-1332333022322220-2113020203133030-0211110101302230-3220111212103022"></a>

<a id="canonical-0231210012313123-3313321210320132-3131330223230032-1122211231323203-2322033223031313-2021213211210212-0101213122000032-2323112333002211"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [custom_hash_algorithms](resources--workload--reference--group-010.md#canonical-1100332230200022-1320301123211133-3301233012201000-2212312020330103-1313303313301322-2011022030012011-3021030200200130-2121002333332023): complete subsection reference.

<a id="canonical-3201312233121311-1312222330331122-3231212203323222-3100311303311231-0330313131011003-0000031121002310-3023220002310322-1211020202210301"></a>

<a id="canonical-3020200213131103-3001011211111011-1313222310013102-2023303103313213-0301203232000223-1003012300033303-2302230031311003-1010231321201101"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-010.md#canonical-1120200202033020-3301221130122030-1333020301210113-2112321001230103-1131212222223230-2211110200221131-2121200222211120-3011112002022220): complete subsection reference.

- [private_key](resources--workload--reference--group-010.md#canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-010.md#canonical-2222002321020303-1031121311231212-3123122030010220-0230112323220000-1200102022001220-3002100212022330-3331312123031020-2101231133103232): complete subsection reference.

<a id="canonical-1202110023222001-3100313121312113-2233103210010221-3332332021103000-2331000312013011-0212033213020210-0112201202301013-0010022021211032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold

<a id="canonical-3030030212212013-1113230132032232-1031303233200223-1223200320201212-0301012320221003-3111210130203133-2321012010031201-3200110312230011"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-0030031130012032-0210202122212233-3211133233221132-0221130312301133-1310210002111222-3312022332223321-0012100322132021-2123120102120013"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold`

<a id="canonical-3303003332311320-3130100232222121-1200301110100303-0212333221020130-0112013113002133-1130323133003010-1133010022030010-2220301002112310"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-2233001022320111-2300032022112012-3312101210213000-1223110132223032-3331003103201130-3000020230020020-2122002013131103-1000121122310123"></a>

<a id="canonical-0201210203202130-2212332003332203-0323010302130100-2000223103133113-2113213323222122-2210323322103313-0211132000022000-2233203101213122"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-3022312220110021-2201321021112121-1213122321300202-0310122230321020-1003021320013213-2332101110133110-2321331232220030-2000232131302220"></a>

<a id="canonical-1332012011332211-2300200133110023-1231313121303323-2113310311112003-2021210011221131-3221333331223032-0222232231203020-3312132020330032"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-1102231223100013-1203322333222203-3332011211023222-0112130120010100-3022012333132213-3030313020231120-3210012122211023-0233100332021121"></a>

<a id="canonical-1220033132223231-2302301220032331-3323031331220030-2301111311301113-0232223230100132-1000223110210013-2113310332003332-2321332302123001"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-2221322203210220-2012031111203011-1013001132210212-3203033330320212-2003113231011313-0003311213030010-2300211123011101-3233223220333200"></a>

<a id="canonical-2300333312323222-2020012333221211-3133013210013123-2013202300002011-3303331032011003-1222200110011333-2210222132321232-3222220303032130"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-2320012301323031-3300313310033100-2023113303332012-3023330301313010-2220232003131022-3303212011113301-3120010031000301-0300220001112001"></a>

<a id="canonical-0032212101012133-2310320033311001-2003032333213330-1121210121330132-3101103201111331-1220131232320223-2131213112302133-3112331032211112"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-0200201301130223-2303001313200302-1221330330100220-0031012233111230-1103320322122333-0033301023113211-1212003103022302-2332112222310330"></a>

<a id="canonical-2012333123213231-3133220300122010-2302103230011200-1020113011231021-1013333120321102-3203031032021202-3230123133003020-2103112220030030"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-3310020123200022-0123011010120221-1113313302332220-0323300202132332-2313301211002331-3133333022133303-3221013003232303-2311330002111320"></a>

<a id="canonical-0312102001333103-2230101221220000-2322121120230321-0221232101312013-2020021202120020-3101032120312032-3320023200200111-3332302121002020"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-3331332210000110-2221300011323222-3010121213232212-0132323110002203-1012321001130102-3300211301222203-3322110103322302-3213000010100012"></a>

<a id="canonical-0210001310000031-0112110132112323-1030023003302102-3103100300222332-2121131213210030-0332333020320332-3003023123212111-2233203000031202"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-3232011032103021-1101321233002311-1331332122010203-0320121312321033-2223200223232030-3133203131023300-0332322303203320-1021331113113230"></a>

<a id="canonical-0033323030131213-1133232121211310-1122323233201100-2103211201330231-3323011003213233-2323322233331113-2020023202201232-1230111301033202"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-0232102302301011-1332110203311101-3131013322000101-0231033011123121-3001230311330323-0213013333133003-0230201022233303-0011230110002320"></a>

<a id="canonical-0233011101103111-1213123110333231-2310222110132121-2031023300103001-3310012122311130-1020101202313300-1211022223330203-3323013201333033"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-1223233330302212-0232321032120223-1212312132201032-2300032322001122-1322300021230003-2103203200122103-2322211203203320-3022022200320331"></a>

<a id="canonical-2221103101003111-2300313203202110-2002112100003132-2100011221112002-0031122303111320-3202001111203202-3200201233320231-0322232302303112"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-0011213123100201-1033103001313301-3310330221203111-1331033203010303-3300113132300123-3201123003011111-1102032000320213-3100023320033203"></a>

<a id="canonical-0302103300120212-1221120133320020-1010232013013122-3003300013022222-3011132302310201-0303103320113102-1223103202213010-1103001000232211"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-3211312231202112-0312202232110111-0111112233111023-1130333023020232-3100303330203303-2301130122133021-2110313002330102-0132022132220300"></a>

<a id="canonical-1311023133122222-1000310020210331-0232232303330310-0032033231321033-1310133011231100-2320333020302003-1103012131002100-2313202323121221"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3201231231313031-0300210322303303-3112101213322223-3013013231023303-0301131330132302-1100103111133132-1000311021302330-3221032111133313"></a>

<a id="canonical-0000333211230302-0010313313222331-2333100323111131-3113202311211110-3323212332211113-3120321111003311-3220023202002313-0121033110312221"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-1211103013113122-0011321220212122-1330332221112123-3003013320113311-3020233313001202-0023123231321033-2122311133220123-1310100323332110"></a>

<a id="canonical-0211201131331123-3133123300320033-1200303023333102-0111323031323012-0101002323333321-3130001111112131-2200301033232032-1001020332100010"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-0100102312122312-3202202133222321-0133031102312003-3301031112330112-3133233303122003-0312001102313101-2321200312032213-1303313210131020"></a>

<a id="canonical-1311212200023100-1020201213331303-0130223131103133-1232110100101221-2330023310102103-2100320022132130-0102113321221023-2333332032310012"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-0220023301023030-0100310113231113-0220301021332312-2023221323330332-2202102102021230-3301300312020122-0021203020232321-1002012321123030"></a>

<a id="canonical-3323313213223013-2021013320310112-3112303102220310-3013032233302131-3212302020300323-2223213333211200-0303022113302132-3033220303001302"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-1202311100232111-2303211131031202-3312012201131213-3321122211103022-2111211123120113-0203112120011123-2102002031030213-2033021210312311"></a>

<a id="canonical-3011002332133201-1010022221320320-1201333022311202-2220110131200002-1211120200023112-0013102031202031-0030220011131030-1232123002112202"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-1100332230200022-1320301123211133-3301233012201000-2212312020330103-1313303313301322-2011022030012011-3021030200200130-2121002333332023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-1201100322031311-1210031320031130-2311210121003131-2102020203201200-0012232131320222-1032231212131200-2133123212102023-1111210230313121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130321020000202-3200233310331000-2220110322221000-1001010003232030-2102233111002233-0123030111000322-1220203130120213-0212000021231002"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-0230312223102321-0200020012102313-1330123232303210-2101130210023032-2132311302221120-1300201112200012-3202011020000331-0131101221000300"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1120200202033020-3301221130122030-1333020301210113-2112321001230103-1131212222223230-2211110200221131-2121200222211120-3011112002022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-1030220020203112-0311031231230022-1302213103223331-0001313221301112-0112021111302031-1211110320220031-1223331302210211-0232032221030011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

Additional upstream details:

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-2113110101032232-2113221003212123-3132210101331223-1311302301300132-2302010020330122-1010330132113003-1332321311222100-2030003311132013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112233031211311-0001311312303111-3133332002131310-1023003022311112-0301122331230211-3222103323023121-1221021030330310-1333220202333113"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](resources--workload--reference--group-010.md#canonical-3220011020230211-2020302110311231-2301203030223111-1003113120302033-2030333022001033-0110001023230212-3100220220312002-1330303033123010): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-010.md#canonical-3233131032002012-1023332003300231-2203300311232200-1002033210211020-3013320211001002-3112231303133232-2033211123101332-2320202022010120): complete subsection reference.

<a id="canonical-3220011020230211-2020302110311231-2301203030223111-1003113120302033-2030333022001033-0110001023230212-3100220220312002-1330303033123010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-010.md#canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0230332201020213-1202020121132212-0123032331322123-2001302031201123-0303033211333221-3311233212330310-0032301301112310-0010210300312031"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011103032320213-2230203300303212-1133310213201313-3310323102100313-2331321102223311-3111111222310301-2110013102212102-2221102130331300"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-1201232123322310-1020131323012130-1323323122312210-0203233331133312-3212323211123221-0200333120303222-1123321220111032-0001323013221323"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2213220003130110-2313200223111100-1122010122032330-2332200110000332-3320201020100012-0110300130202320-3022231220011021-3323322220111112"></a>

<a id="canonical-0112223133110330-1133203012100322-2022233212303311-3200222011303212-3101231321323020-3113320221123001-0222330330321130-3332300321122103"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2013110012233133-0300122021330213-1020131100122002-1320020303122133-0022230330001020-2203030031211312-3303001331301032-2323002021230131"></a>

<a id="canonical-3011310332123200-2211002122232033-0330331222313211-3023103220011100-3231313121213210-2011001131303010-2113033003312331-2123323201013031"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3233131032002012-1023332003300231-2203300311232200-1002033210211020-3013320211001002-3112231303133232-2033211123101332-2320202022010120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-010.md#canonical-0000103021020013-2333222020321111-2221322222122222-1012301222211232-3331320010202111-1101320110232221-1222102310300230-3123220221113213)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-0022223303312122-1201000322010330-1000333312112310-3233230113320233-3222220010230201-3312000301330230-1321003131333031-1003121321021212"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321020000223333-1230220030113122-3201331202013231-0211002021300200-1300131330323220-1223220311332313-1120132313211312-2113233313230103"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3301320302110203-3032323200013001-1121032012303003-0301211332201003-0031120020332000-2230212123012231-2001310323310301-3100212001130230"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2033232020213113-3332123103201231-3112211101031312-0332301330121213-1111313132133203-1220123320232010-0230223330323231-3122322213322230"></a>

<a id="canonical-3123322223312203-1021030220001012-0001123211301201-1121330030321012-3030211200002023-3321123203320131-1211030201120212-1233210303030131"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2222002321020303-1031121311231212-3123122030010220-0230112323220000-1200102022001220-3002100212022330-3331312123031020-2101231133103232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-010.md#canonical-0220313011200113-2210220130122300-1020333311303231-0221222302220012-3120013102133220-2132302001202230-2322133133130313-1313320213312221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-1000220113001030-0031103113122231-2110003022203301-3222221222302210-2101133210332011-3330113010321002-1101230212102000-1222312202230130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

Additional upstream details:

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-1021112010022003-1232212023220022-1131231223030003-3001312101202103-0020312131033210-1201003023023020-2311332101003112-0213010211311223"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001203233221300-2321321002223122-3211002031201202-1022002330220201-0311101123211321-3120012202131120-0310001213012130-0013022313101120"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config`

- [custom_security](resources--workload--reference--group-010.md#canonical-1231032221230000-1010100200221203-1121220011012313-3332330220020233-1100222130021303-0203321102313101-0211023223331023-0011323212302011): complete subsection reference.

- [default_security](resources--workload--reference--group-010.md#canonical-2212231013132322-1113323133200030-3133300033021130-0100211301333112-3133022222320232-3120322131131110-3113011032331223-3312133132103033): complete subsection reference.

- [low_security](resources--workload--reference--group-010.md#canonical-3230331013322123-1233332021013212-2301010301100300-2130203133232103-2300113022202231-3013013233020330-0323212301221010-3003032120203312): complete subsection reference.

- [medium_security](resources--workload--reference--group-010.md#canonical-3221312221332112-3133132000110333-0021212030322333-1222322221303102-2111122201230103-3320102333230330-2101323311022232-0032302321133232): complete subsection reference.

<a id="canonical-1231032221230000-1010100200221203-1121220011012313-3332330220020233-1100222130021303-0203321102313101-0211023223331023-0011323212302011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-0001022220131201-0012010223003002-0221120022222010-0123320300102022-3333213113311320-2300322103310123-1111331333032030-3033012030003302"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332033010032010-1223030230102022-2223100313021220-3313213311103023-2232313033000123-2131123033323020-1020012213332303-2211220110032001"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security`

<a id="canonical-1220130030100132-2211000312211001-1022020203120303-2303331110021232-1130232323231123-1023131032100122-3303011001120332-1031322313223322"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0031122200122322-1113230120223003-2012010012000223-0322300032010113-0312201130122301-2201311123100002-0111112201233332-2123123112111122"></a>

<a id="canonical-3303300121113103-2022131022303001-3231321111333111-3212331202023231-2230032312030333-2311233000223203-2132011031020130-3123033111202302"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-0030220231123331-2301102210220311-3001033321122131-2131300100302211-2323133020322010-1130320110131120-0301332332311303-2331011310120201"></a>

<a id="canonical-0211032012300033-2331221231303131-0002023230321033-2031122112130201-0020022122310222-1101100021200322-2022232010121003-0223322003233310"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-2212231013132322-1113323133200030-3133300033021130-0100211301333112-3133022222320232-3120322131131110-3113011032331223-3312133132103033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-2223030220201032-3013130030311211-2113222113102022-0222001001021223-2330130220300120-3313303131230030-3200232010321112-2020032122302320"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230331013322123-1233332021013212-2301010301100300-2130203133232103-2300113022202231-3013013233020330-0323212301221010-3003032120203312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-2212131332110032-2020021232220111-2032120221213200-1020103203031131-0022311001332201-2113113111130202-3230212233210230-3333223313022301"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
low_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221312221332112-3133132000110333-0021212030322333-1222322221303102-2111122201230103-3320102333230330-2101323311022232-0032302321133232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-1233203302211212-3300120033202101-0200330001312321-1222133023031223-1233313002123211-3203212322123012-1131031112003021-1331232103023221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-1332123033032010-1110111113002120-2121032330331233-0231330000322133-3011120200122021-0012301213201331-0122331203100210-1003120123132312"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
medium_security = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-2120111023132302-1210312111201023-1030011013200013-2033120121331120-0021231020003133-2022012201133213-3121202030132031-2101000020232022"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321302212210203-1311333302222202-1313221013001101-2232013232302232-1221311201301321-2013021033030122-2322001211210211-1113022202223222"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls`

<a id="canonical-3103233303331132-3313212021122001-2120312023233322-2233212330310132-1222001011322332-3230211232120112-1022100200030211-3210023012321102"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.client_certificate_optional` property

Type: `"bool"`. Optional.

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

- [crl](resources--workload--reference--group-010.md#canonical-2001333030330320-2023002103031001-3002232311311032-3022320130133223-2220003002311012-2200213221110010-2011130013200033-3312301120331202): complete subsection reference.

- [no_crl](resources--workload--reference--group-010.md#canonical-1311102021032031-3333223133333320-1000312233313312-1201011122102200-2120123032230012-2121221321023031-2121231012212103-3112303201103323): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-010.md#canonical-1320023331013030-2310123101313321-2010323333101333-1131332221202102-1130032102133222-2111333301023022-0032001012303232-0102013111132220): complete subsection reference.

<a id="canonical-3302021222231310-0012120021310322-2231202120121221-0130002022333320-2021331213313133-1333002303230003-1330032010330201-3111023010021133"></a>

<a id="canonical-1111001303302331-0221020201113320-0303031221312211-0302202102321030-2133020220331011-1233220133300332-0313103222120010-3113323003030202"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca_url` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [xfcc_disabled](resources--workload--reference--group-010.md#canonical-1123230303002221-0101211213322212-1222302011111321-1110212023133202-0122011213310112-0310310110213111-2023030303012010-1310313101030302): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-010.md#canonical-1022001030312122-0110021131331203-1321111312211223-0231321110200201-2232010002033222-1321002220133122-2012110202313210-2130032010303332): complete subsection reference.

<a id="canonical-2001333030330320-2023002103031001-3002232311311032-3022320130133223-2220003002311012-2200213221110010-2011130013200033-3312301120331202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-2130211313312130-1210132332123120-0122213002122302-3031021020233210-0102033001112103-2113103031332330-2102000133101013-3020033001011101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232133120222131-2222211001211133-2011230220131222-1220110303322100-3113031202110202-3010330322110333-3313023221010321-2022103130330230"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl`

<a id="canonical-0022102312301022-0203101333013233-0320020132320212-2100030112033000-2021000100303330-0100023102213030-0221212001012333-3333323101120222"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3322203013220313-0120313100003113-3222233102011012-2110013131100303-0012133222332101-0023323230211022-2311123202023103-0200033320000113"></a>

<a id="canonical-2310233220231321-1110332011200312-1130003200020223-2123113121030033-3210030022132303-0012112102003132-3001023110013103-3011210021113002"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3222222133021332-0222232322230332-0030001310233100-3231022000220032-0312321103123202-1210233120300110-0030212131302202-0130333302011010"></a>

<a id="canonical-1113110210010230-1303112000002031-3103311022202201-0120102312021003-0321303233321111-1010333000220011-0200300223300121-2320023031132103"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1311102021032031-3333223133333320-1000312233313312-1201011122102200-2120123032230012-2121221321023031-2121231012212103-3112303201103323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-2123102233120113-3211220313301001-1211211111130202-1032110131233333-2123300202023322-0220110020212130-2221323233002232-0311013320202032"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
no_crl = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320023331013030-2310123101313321-2010323333101333-1131332221202102-1130032102133222-2111333301023022-0032001012303232-0102013111132220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-2010321033203023-0122011110100121-1221203302303013-3102222030211032-3030202121203130-0122110101310123-3222230333023203-3331213210023300"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310122101120013-2102323003313100-3032331213220020-0201011120131100-2322031133012022-3300332332323221-3023313111011102-0233213320013133"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-3230313123330033-2123122201133011-0233321322211220-1320333013023100-2102111330103103-2310232112221100-3013013300220212-2002100002020130"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2301222132203210-3313223021320213-1333012323102103-2321110102103231-1220011010000011-2102232120202323-2211302320121130-2223301103311000"></a>

<a id="canonical-0023312012012333-2323101002212131-0213300030200123-0200313001021302-3132103000032310-1221331321202031-3012331333232122-3112332222233211"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1032123320321033-0200312331001112-2210300102020101-0103021003000132-0330023033121112-0131332331122102-1321233223133001-0203221212303122"></a>

<a id="canonical-2312211211301130-0322112211322030-1331030111322120-2311313030000102-2303013333212020-0131003313222212-1120211130132221-3221211022312121"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1123230303002221-0101211213322212-1222302011111321-1110212023133202-0122011213310112-0310310110213111-2023030303012010-1310313101030302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-3111330312223100-3212113002300021-3011312000233330-2033302013303200-3130213332312211-0221032030232231-3311213300212120-0332001223203000"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
xfcc_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022001030312122-0110021131331203-1321111312211223-0231321110200201-2232010002033222-1321002220133122-2012110202313210-2130032010303332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-009.md#canonical-3212310132100213-2000301220120030-0112323230310220-3200032133332210-3223310210320320-2331013310222010-2010102133130322-2111123310321223)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-3230132323111312-0132020110123203-0222302101120231-0202012321012303-1213310020220000-0101203211110131-1303320330302033-3001313221002112)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-3310220202003120-0020232310001311-1230011321302031-3002200220020220-1002102031303233-0322332031301112-0031333023233103-0133102211311201)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-3030102300301111-2132021123223002-2331131033102101-0113203222021332-2122131110023202-3302330233311223-1133210311301333-0202330000301013"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023103223102031-0302023011123113-0102310230221103-1200300312130123-3230220302122220-0300223320112130-0311030030301122-1331003230323133"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-3021102022303222-3100111111023010-1211032231033322-3303321133211131-2113102122223303-2001111131202233-2120211303313320-3020132200313320"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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

<a id="canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert

<a id="canonical-1232112013302333-0303103031303231-1222031113202031-3331003311332113-0013301133323011-0302111121301200-1222300233002111-1220032321103222"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110321121232120-3321012331231232-1221301112323021-1322110330200133-1302112320200223-2030312221011202-1100310303112130-2000332331123331"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert`

<a id="canonical-1323133111232320-3112232100031121-2202021111201223-2133333211101030-2131122303022111-0133312013012322-0101233220003322-1023101020331222"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.add_hsts` property

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-3123302120010331-1230221202032032-1212012102323313-0012100223331031-3133022231220003-1233020310102221-3302210300301033-0021230202220333"></a>

<a id="canonical-0300031012231133-2030221103321213-2110200132332111-3022323322212321-3121222123021120-2010001320103002-2330010233122300-3020332003321012"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.append_server_name` property

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--workload--reference--group-010.md#canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232): complete subsection reference.

<a id="canonical-3132312010312213-2323033210000301-0012312330222002-0301031213332012-2211201112113331-0311231122310110-0302303011200231-2012112203222211"></a>

<a id="canonical-1022103211231020-1220021230020200-0210132111310031-2331031021102330-0132310320330221-3333313133333010-2210011310213012-2300002012131321"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.connection_idle_timeout` property

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--workload--reference--group-010.md#canonical-2212010221300233-2201101023002210-3010002022323233-3113011212030131-3032203233302230-0231011311211003-2200023023000231-0320202103113033): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-010.md#canonical-0331230213133123-1130310110210111-1321130212223101-0320232323122321-3111200212112300-3031202031023212-0122001001002002-3111023022301313): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-010.md#canonical-0123302220301321-0132211023131201-3220202102311211-2131030202102302-0221310222332302-3321210201332223-3210110121201203-0011033011213313): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-010.md#canonical-2301110231331322-3310221221300232-1201002123223112-0122032021031011-0101203313332122-1312001131032013-1101002312122203-1030033232131121): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001): complete subsection reference.

<a id="canonical-3112313333123023-0132221002100312-1023212222003312-3132111201200113-1311223332010320-0221013032333202-2012300020320102-2320003101330223"></a>

<a id="canonical-1030213122030302-1322000300101220-0310301311103201-0311220201211230-2221300131101000-2131212131203303-1022323032210201-3321000330302323"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_redirect` property

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

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

- [no_mtls](resources--workload--reference--group-010.md#canonical-1222331122321301-0002020313232131-2111200221222211-0133120203310331-3111211101033101-3212020231221033-2003023203022002-0202201133311302): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-010.md#canonical-2212121030220022-1122202123012230-2320133132212021-1010311121113312-3023212110310032-2101332003123313-0112003011103233-1110123011301312): complete subsection reference.

- [pass_through](resources--workload--reference--group-010.md#canonical-0011102022130121-2312331112021323-3112133222310201-1022122230332333-1001102130231330-0303221022113011-0131100333200021-3212121001001113): complete subsection reference.

<a id="canonical-2312202332200003-2311111301131023-1102300300300101-3300201213203023-2111201211122230-0311033133113211-1201121032223203-2311221110121113"></a>

<a id="canonical-0232313132020100-0100132233321010-1002012223332021-1101220230133230-3121201030002023-0003332012202010-1103011122121012-3001013303032033"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1212100333110003-0010030003023213-1331133201222211-0023101302331222-2113003102130220-1211023021321100-1223122303012002-2221310200100222"></a>

<a id="canonical-2030300301003203-3103311001120031-2303320201012331-0011211001022013-0013210121110012-1000220110203021-0023301113212213-3210012200011202"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-1010030321220303-1302202103201122-3302102232200220-1322202103121200-2303220033323201-1131100210033120-1123130023323033-1212311121112030"></a>

<a id="canonical-2002030231220330-1203111033322110-1011101320032303-0300131021032031-0001131301322311-1110023012130210-3111300233131220-2112110301331112"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.server_name` property

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](resources--workload--reference--group-010.md#canonical-3113321322213012-1133311332223211-0212301313120320-2200201201212023-3113301313231213-2111031111233222-3330001032333212-3112113321000021): complete subsection reference.

- [use_mtls](resources--workload--reference--group-011.md#canonical-1112020130200033-1212030123322311-1103211123233210-3303131123200003-0033323320312331-1303132312122133-3032112313313232-1220000112111220): complete subsection reference.

<a id="canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-0032231120300313-2312012110303100-0303011330133331-2312133233023020-1302212330033032-1010202323322031-0202330220101333-2223023023312000"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301010322200001-2132031111023300-3120222231010330-0210332100201033-1222211203032211-1321201232200231-2311003000102333-3210133003000130"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options`

- [default_coalescing](resources--workload--reference--group-010.md#canonical-3001201313233020-0031113321003221-1223021203033001-0303222112103221-2020323230012131-1120012022332202-2033301103122200-3322012122333032): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-010.md#canonical-3001101301333012-1001032231123013-0022102001110211-2121001331021222-2121200030001103-3010003322331333-3202230111110223-0131101301031023): complete subsection reference.

<a id="canonical-3001201313233020-0031113321003221-1223021203033001-0303222112103221-2020323230012131-1120012022332202-2033301103122200-3322012122333032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-010.md#canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-1113301221232113-0231122113222320-3111202102101023-2202023233220123-3301331020333033-0213131031312232-1000031310013100-0323313131222200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

Additional upstream details:

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

Terraform syntax:

```terraform
default_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001101301333012-1001032231123013-0022102001110211-2121001331021222-2121200030001103-3010003322331333-3202230111110223-0131101301031023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-010.md#canonical-3210303103310002-2310113031123222-2101231231122213-2233021130122303-0233321001201002-0330322101202322-2003321331113003-0022133312300232)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3122312100033321-3122201130333033-2002021320033212-3030003021123132-2211100012203200-1221200031113113-3110022032332232-1203313100200120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

Additional upstream details:

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212010221300233-2201101023002210-3010002022323233-3113011212030131-3032203233302230-0231011311211003-2200023023000231-0320202103113033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-2103110332300320-1130103030222311-3000012332103200-3003311111112302-0132003013003102-1321331322123002-0122231313232121-3331000003333123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

Additional upstream details:

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

Terraform syntax:

```terraform
default_header = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331230213133123-1130310110210111-1321130212223101-0320232323122321-3111200212112300-3031202031023212-0122001001002002-3111023022301313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-1101130120102203-1021132231003031-0031100212231322-1302322322203203-2212220310311332-0311033330211021-3120301012231111-0123122231201230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

Additional upstream details:

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

Terraform syntax:

```terraform
default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123302220301321-0132211023131201-3220202102311211-2131030202102302-0221310222332302-3321210201332223-3210110121201203-0011033011213313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-2313312033200311-2020223013103232-3320231001211002-1200220310100110-0213221001030311-2322021201201020-3021221020012103-0233331202231021"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
disable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301110231331322-3310221221300232-1201002123223112-0122032021031011-0101203313332122-1312001131032013-1101002312122203-1030033232131121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-2320121220311111-3333022012222323-3200311102331320-0302212003020121-3123300021110301-3131130203321233-1110232010321323-2202110223300331"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
enable_path_normalize = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-1223310002012002-0110033021102331-3310330013101112-0002100221312013-0322312333233032-3013230310211102-0331132200301012-3212123311222202"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123113233210311-1111231200031012-2121220211312322-1211310332331021-2103110103213001-3202222132332230-3310031333220301-3012222122332231"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](resources--workload--reference--group-010.md#canonical-3131320122021222-1122202213031111-0311122112320212-1002111302121300-1312200303033032-1020201110003003-0211332003112122-2113212112233200): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-010.md#canonical-1223222121310230-1300021101211033-1132021133012330-1010020023021121-1313030322131100-1023321112331222-2231313112223331-3321101313100120): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-010.md#canonical-3030001013313220-3110300130210113-1031003122323110-3321330320103001-3020233032232131-1322033202000033-3133030023223232-2200133133102122): complete subsection reference.

<a id="canonical-3131320122021222-1122202213031111-0311122112320212-1002111302121300-1312200303033032-1020201110003003-0211332003112122-2113212112233200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3031011122312232-2012312121130212-3101032233223213-1331130003232033-0002123313131213-2301333002300221-1213110232212110-2201011131202303"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223330003120302-0020313323231032-1201230102121230-0320203012200000-2111012131102100-2220333300033132-2131121211021230-2333301322102021"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](resources--workload--reference--group-010.md#canonical-3010113201320022-0210032323222211-1211322130103300-0000312020130232-1200030233100102-0130112013113031-1311320113310023-0331023032211221): complete subsection reference.

<a id="canonical-3010113201320022-0210032323222211-1211322130103300-0000312020130232-1200030233100102-0130112013113031-1311320113310023-0331023032211221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-010.md#canonical-3131320122021222-1122202213031111-0311122112320212-1002111302121300-1312200303033032-1020201110003003-0211332003112122-2113212112233200)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-3133130033301212-2333110311100113-0013230032310320-1101120010212111-3233210032123231-1123303331201333-0001022320220310-3130130020231013"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222131011330033-3333213311221331-1030312123203013-3203111330220120-2020322230001212-2003022011113210-1221102103323113-1110120210233322"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](resources--workload--reference--group-010.md#canonical-2230300033123123-1321223203201102-1233111113113111-3121201001112100-3330100132321203-0133321120111210-2022032222123202-2223323030112222): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-010.md#canonical-1132012212031303-2232323312230123-3230320121033303-2202031131103230-3121210333220122-0123103323131010-2200332011123030-1223102100031010): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-010.md#canonical-0301110313233103-2030032023222201-3303001310100120-0211220003033313-3010022332231133-3131213330113313-3131232322223233-2111132103323103): complete subsection reference.

<a id="canonical-2230300033123123-1321223203201102-1233111113113111-3121201001112100-3330100132321203-0133321120111210-2022032222123202-2223323030112222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-010.md#canonical-3131320122021222-1122202213031111-0311122112320212-1002111302121300-1312200303033032-1020201110003003-0211332003112122-2113212112233200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-010.md#canonical-3010113201320022-0210032323222211-1211322130103300-0000312020130232-1200030233100102-0130112013113031-1311320113310023-0331023032211221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1103301123323323-0003230312310212-0132130110030211-0322201130232221-1201012132223021-3112322231003212-0220110320111132-1021212130033022"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

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
default_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132012212031303-2232323312230123-3230320121033303-2202031131103230-3121210333220122-0123103323131010-2200332011123030-1223102100031010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-010.md#canonical-3131320122021222-1122202213031111-0311122112320212-1002111302121300-1312200303033032-1020201110003003-0211332003112122-2113212112233200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-010.md#canonical-3010113201320022-0210032323222211-1211322130103300-0000312020130232-1200030233100102-0130112013113031-1311320113310023-0331023032211221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2223323120032013-2223311013322001-0301001002120022-0230312033212033-1313030221130212-2102023001230221-3010033203313032-0010002133002112"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301110313233103-2030032023222201-3303001310100120-0211220003033313-3010022332231133-3131213330113313-3131232322223233-2111132103323103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-010.md#canonical-3131320122021222-1122202213031111-0311122112320212-1002111302121300-1312200303033032-1020201110003003-0211332003112122-2113212112233200)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-010.md#canonical-3010113201320022-0210032323222211-1211322130103300-0000312020130232-1200030233100102-0130112013113031-1311320113310023-0331023032211221)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0022202132222030-2011203331303101-1011013303003310-0001021321021031-2203012203320133-0232221223213212-0123120011133223-2113330102313211"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

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
proper_case_header_transformation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223222121310230-1300021101211033-1132021133012330-1010020023021121-1313030322131100-1023321112331222-2231313112223331-3321101313100120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2333231320123222-0331033333101232-3120211032030012-2211033203101102-0301201230002230-3033130221301230-3012311233312100-1310021212200002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

Additional upstream details:

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

Terraform syntax:

```terraform
http_protocol_enable_v1_v2 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030001013313220-3110300130210113-1031003122323110-3321330320103001-3020233032232131-1322033202000033-3133030023223232-2200133133102122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-010.md#canonical-3321233023232103-1212213330303210-1111113130313312-1323223120210100-2202312311101302-1303212320303032-3030011002101032-0231021313323001)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0113311002103100-3101323221202232-2031321310213213-1321110020322011-3313331013102131-0202233000023101-0213033021021003-3321331102300103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

Additional upstream details:

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

Terraform syntax:

```terraform
http_protocol_enable_v2_only = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222331122321301-0002020313232131-2111200221222211-0133120203310331-3111211101033101-3212020231221033-2003023203022002-0202201133311302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-2112232222201122-1220320210102222-1003202002011233-0231321321222001-2130020221133302-2112333213233322-1103231132200023-2112312300012023"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
no_mtls = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212121030220022-1122202123012230-2320133132212021-1010311121113312-3023212110310032-2101332003123313-0112003011103233-1110123011301312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-3223010232333120-2322332130001330-1122101120220111-3022332021212112-2032201200302202-0102323102120001-3003222322313233-1310200103012020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

Additional upstream details:

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

Terraform syntax:

```terraform
non_default_loadbalancer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011102022130121-2312331112021323-3112133222310201-1022122230332333-1001102130231330-0303221022113011-0131100333200021-3212121001001113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-3323111003031111-1203021301333003-0101230232103330-2310003003301111-0222031203130330-3013031222222332-0331223012123322-2121110020013000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

Additional upstream details:

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

Terraform syntax:

```terraform
pass_through = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113321322213012-1133311332223211-0212301313120320-2200201201212023-3113301313231213-2111031111233222-3330001032333212-3112113321000021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-1232202132323330-3121203211322020-1003300111202323-0111311021333131-3212021023322302-3020022221133230-0222331213203003-2011233113202121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110330031003333-2003331020132031-0133322231310310-1322310003300123-1233200001101202-3002020021113022-3333220032213320-2303103012311012"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config`

- [custom_security](resources--workload--reference--group-010.md#canonical-0102100311133102-0002303333230031-1302323211132313-1332010113331202-3211331330333012-1220333313200011-3221231330220033-2312031210103030): complete subsection reference.

- [default_security](resources--workload--reference--group-010.md#canonical-0130322011203020-3110233320032100-2310033122101022-1201120110211331-0320032230332201-0012000310213022-1010333320302011-3201333103311323): complete subsection reference.

- [low_security](resources--workload--reference--group-011.md#canonical-1002310211223003-1200302131331023-3131123210310123-1222310013023103-3223211302133110-1203321311233321-2323021312211301-2131213022221011): complete subsection reference.

- [medium_security](resources--workload--reference--group-011.md#canonical-0333222112313030-3113012230130130-1220032020010012-2212210012213021-0120213112220203-1132022303133230-0320033231233313-0113123132132221): complete subsection reference.

<a id="canonical-0102100311133102-0002303333230031-1302323211132313-1332010113331202-3211331330333012-1220333313200011-3221231330220033-2312031210103030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-010.md#canonical-3113321322213012-1133311332223211-0212301313120320-2200201201212023-3113301313231213-2111031111233222-3330001032333212-3112113321000021)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-0311103232230122-2300302333111101-2121300321122122-2000010102101100-1201010022130210-3132302331331022-3033103121022032-2020133211001030"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013101013202101-1221312031323230-2101203010230220-3333011130010012-1011030132202000-0331000121130323-2113031002123312-0021003002300322"></a>

### Direct properties for `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security`

<a id="canonical-0033103303111230-1030302332032333-2213000011311321-2102121232130310-1323311121100120-3201233101322113-1101022112110030-2113223321210223"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1031322302330232-1302030312201002-3110302132103010-1112330310110132-0001302030300130-2020212033231222-0331112110302302-2011321302131233"></a>

<a id="canonical-0132022222032221-1211111030220330-3311230222221322-0223233222021303-3010010030122030-1120222313122130-2303312031021221-0210213223201211"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.max_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-3202112000200121-3220200312110202-1030030011011322-1011022320123133-3333020123111131-3133100131020320-0112101032212323-1002322320002331"></a>

<a id="canonical-2003213103112021-0030331002221221-3030211000020331-3223002020301323-1102012012132232-2110233123032302-2030002123131003-1201011320300302"></a>

#### `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security.min_version` property

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

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

<a id="canonical-0130322011203020-3110233320032100-2310033122101022-1201120110211331-0320032230332201-0012000310213022-1010333320302011-3201333103311323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-006.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-006.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-009.md#canonical-1222100102223033-2323000102001021-0112111031020311-0221301000203113-2112023321023331-2013201202121333-3010012033322311-0310303301113101)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-009.md#canonical-3210102113031103-2332201101220100-1103122221002212-3002101012001000-1332031130002011-0021231133303122-2121222212212213-1123320220010203)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-009.md#canonical-2311132233010120-0223230230122210-0332200203312210-3000132111100333-1133101301233010-3211310033122333-2323002110210000-1300012002110122)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-009.md#canonical-1202020123020333-2332300320113330-1321112113331202-2112220001121101-3022001013310321-0020100001212100-0032232320012323-1311122212310031)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-0123203011200310-2323132333110331-0320211012002122-2323013220112302-2000030132131023-1312320130303213-2223130103233300-3130200202200302)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-010.md#canonical-3113321322213012-1133311332223211-0212301313120320-2200201201212023-3113301313231213-2111031111233222-3330001032333212-3112113321000021)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-1230302200013111-1101103020222130-3202011121022011-1230033131211322-3103332221301121-0010212201303131-2213103013203200-1221032231110110"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
default_security = {}
```

This is an empty object or choice marker. It has no direct properties.
