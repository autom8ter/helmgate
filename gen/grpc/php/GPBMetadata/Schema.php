<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace GPBMetadata;

class Schema
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Api\Annotations::initOnce();
        \GPBMetadata\Google\Protobuf\Struct::initOnce();
        \GPBMetadata\Google\Protobuf\Timestamp::initOnce();
        \GPBMetadata\Google\Protobuf\Any::initOnce();
        \GPBMetadata\Google\Protobuf\GPBEmpty::initOnce();
        \GPBMetadata\GithubCom\Mwitkow\GoProtoValidators\Validator::initOnce();
        $pool->internalAddGeneratedFile(hex2bin(
            "0ae5140a0c736368656d612e70726f746f120968656c6d50726f78791a1c" .
            "676f6f676c652f70726f746f6275662f7374727563742e70726f746f1a1f" .
            "676f6f676c652f70726f746f6275662f74696d657374616d702e70726f74" .
            "6f1a19676f6f676c652f70726f746f6275662f616e792e70726f746f1a1b" .
            "676f6f676c652f70726f746f6275662f656d7074792e70726f746f1a3667" .
            "69746875622e636f6d2f6d7769746b6f772f676f2d70726f746f2d76616c" .
            "696461746f72732f76616c696461746f722e70726f746f22760a0a446570" .
            "656e64656e6379121f0a0563686172741801200128094210e2df1f0c0a0a" .
            "5e2e7b312c3232357d2412210a0776657273696f6e1802200128094210e2" .
            "df1f0c0a0a5e2e7b312c3232357d2412240a0a7265706f7369746f727918" .
            "03200128094210e2df1f0c0a0a5e2e7b312c3232357d24224d0a0a4d6169" .
            "6e7461696e6572121e0a046e616d651801200128094210e2df1f0c0a0a5e" .
            "2e7b312c3232357d24121f0a05656d61696c1802200128094210e2df1f0c" .
            "0a0a5e2e7b312c3232357d24223c0a0b436861727446696c746572121e0a" .
            "047465726d1801200128094210e2df1f0c0a0a5e2e7b312c3232357d2412" .
            "0d0a05726567657818022001280822dc020a054368617274121e0a046e61" .
            "6d651801200128094210e2df1f0c0a0a5e2e7b312c3232357d24120c0a04" .
            "686f6d6518022001280912130a0b6465736372697074696f6e1803200128" .
            "09120f0a0776657273696f6e180420012809120f0a07736f757263657318" .
            "052003280912100a086b6579776f726473180620032809120c0a0469636f" .
            "6e18072001280912120a0a64657072656361746564180820012808122b0a" .
            "0c646570656e64656e6369657318092003280b32152e68656c6d50726f78" .
            "792e446570656e64656e6379122a0a0b6d61696e7461696e657273180a20" .
            "03280b32152e68656c6d50726f78792e4d61696e7461696e657212300a08" .
            "6d65746164617461180b2003280b321e2e68656c6d50726f78792e436861" .
            "72742e4d65746164617461456e7472791a2f0a0d4d65746164617461456e" .
            "747279120b0a036b6579180120012809120d0a0576616c75651802200128" .
            "093a023801222a0a0643686172747312200a066368617274731801200328" .
            "0b32102e68656c6d50726f78792e43686172742290010a03417070121e0a" .
            "046e616d651801200128094210e2df1f0c0a0a5e2e7b312c3232357d2412" .
            "230a096e616d6573706163651802200128094210e2df1f0c0a0a5e2e7b31" .
            "2c3232357d2412230a0772656c6561736518052001280b32122e68656c6d" .
            "50726f78792e52656c65617365121f0a05636861727418142001280b3210" .
            "2e68656c6d50726f78792e436861727422240a0441707073121c0a046170" .
            "707318012003280b320e2e68656c6d50726f78792e417070224f0a094170" .
            "7046696c74657212110a096e616d65737061636518012001280912100a08" .
            "73656c6563746f72180220012809120d0a056c696d697418032001280d12" .
            "0e0a066f666673657418042001280d22a2010a0752656c65617365120f0a" .
            "0776657273696f6e18012001280d12270a06636f6e66696718022001280b" .
            "32172e676f6f676c652e70726f746f6275662e537472756374120d0a056e" .
            "6f74657318032001280912130a0b6465736372697074696f6e1804200128" .
            "09120e0a0673746174757318052001280912290a0a74696d657374616d70" .
            "7318062001280b32152e68656c6d50726f78792e54696d657374616d7073" .
            "2293010a0a54696d657374616d7073122b0a076372656174656418012001" .
            "280b321a2e676f6f676c652e70726f746f6275662e54696d657374616d70" .
            "122b0a077570646174656418022001280b321a2e676f6f676c652e70726f" .
            "746f6275662e54696d657374616d70122b0a0764656c6574656418032001" .
            "280b321a2e676f6f676c652e70726f746f6275662e54696d657374616d70" .
            "224d0a0641707052656612230a096e616d65737061636518012001280942" .
            "10e2df1f0c0a0a5e2e7b312c3232357d24121e0a046e616d651802200128" .
            "094210e2df1f0c0a0a5e2e7b312c3232357d2422dc010a08417070496e70" .
            "757412230a096e616d6573706163651801200128094210e2df1f0c0a0a5e" .
            "2e7b312c3232357d24121f0a0563686172741802200128094210e2df1f0c" .
            "0a0a5e2e7b312c3232357d2412220a086170705f6e616d65180320012809" .
            "4210e2df1f0c0a0a5e2e7b312c3232357d2412370a06636f6e6669671804" .
            "2003280b321f2e68656c6d50726f78792e417070496e7075742e436f6e66" .
            "6967456e7472794206e2df1f0220011a2d0a0b436f6e666967456e747279" .
            "120b0a036b6579180120012809120d0a0576616c75651802200128093a02" .
            "3801222e0a0c4e616d657370616365526566121e0a046e616d6518012001" .
            "28094210e2df1f0c0a0a5e2e7b312c3232357d24223c0a0d4e616d657370" .
            "61636552656673122b0a0a6e616d6573706163657318012003280b32172e" .
            "68656c6d50726f78792e4e616d65737061636552656622460a0d48697374" .
            "6f727946696c74657212260a0372656618012001280b32112e68656c6d50" .
            "726f78792e4170705265664206e2df1f022001120d0a056c696d69741802" .
            "2001280d32ca050a1048656c6d50726f78795365727669636512460a0647" .
            "657441707012112e68656c6d50726f78792e4170705265661a0e2e68656c" .
            "6d50726f78792e417070221982d3e493021312112f617070732f7b6e616d" .
            "6573706163657d12690a0a476574486973746f727912182e68656c6d5072" .
            "6f78792e486973746f727946696c7465721a0f2e68656c6d50726f78792e" .
            "41707073223082d3e493022a12282f617070732f7b7265662e6e616d6573" .
            "706163657d2f7b7265662e6e616d657d2f686973746f7279124e0a0a5365" .
            "617263684170707312142e68656c6d50726f78792e41707046696c746572" .
            "1a0f2e68656c6d50726f78792e41707073221982d3e493021312112f6170" .
            "70732f7b6e616d6573706163657d125b0a0c556e696e7374616c6c417070" .
            "12112e68656c6d50726f78792e4170705265661a162e676f6f676c652e70" .
            "726f746f6275662e456d707479222082d3e493021a2a182f617070732f7b" .
            "6e616d6573706163657d2f7b6e616d657d125e0a0b526f6c6c6261636b41" .
            "707012112e68656c6d50726f78792e4170705265661a0e2e68656c6d5072" .
            "6f78792e417070222c82d3e493022622212f617070732f7b6e616d657370" .
            "6163657d2f7b6e616d657d2f726f6c6c6261636b3a012a124f0a0a496e73" .
            "74616c6c41707012132e68656c6d50726f78792e417070496e7075741a0e" .
            "2e68656c6d50726f78792e417070221c82d3e493021622112f617070732f" .
            "7b6e616d6573706163657d3a012a12590a0955706461746541707012132e" .
            "68656c6d50726f78792e417070496e7075741a0e2e68656c6d50726f7879" .
            "2e417070222782d3e49302211a1c2f617070732f7b6e616d657370616365" .
            "7d2f7b6170705f6e616d657d3a012a124a0a0c5365617263684368617274" .
            "7312162e68656c6d50726f78792e436861727446696c7465721a112e6865" .
            "6c6d50726f78792e436861727473220f82d3e493020912072f6368617274" .
            "73420d5a0b68656c6d50726f78797062620670726f746f33"
        ));

        static::$is_initialized = true;
    }
}

