<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace MeshelmProxy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>helmProxy.RequestAuthentication</code>
 */
class RequestAuthentication extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .helmProxy.JWTRule rules = 1;</code>
     */
    private $rules;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \MeshelmProxy\JWTRule[]|\Google\Protobuf\Internal\RepeatedField $rules
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .helmProxy.JWTRule rules = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getRules()
    {
        return $this->rules;
    }

    /**
     * Generated from protobuf field <code>repeated .helmProxy.JWTRule rules = 1;</code>
     * @param \MeshelmProxy\JWTRule[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setRules($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \MeshelmProxy\JWTRule::class);
        $this->rules = $arr;

        return $this;
    }

}

