<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace MeshelmProxy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>helmProxy.AuthnPolicy</code>
 */
class AuthnPolicy extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.helmProxy.AuthAction action = 1;</code>
     */
    private $action = 0;
    /**
     * Generated from protobuf field <code>repeated .helmProxy.AuthnRule rules = 2;</code>
     */
    private $rules;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $action
     *     @type \MeshelmProxy\AuthnRule[]|\Google\Protobuf\Internal\RepeatedField $rules
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.helmProxy.AuthAction action = 1;</code>
     * @return int
     */
    public function getAction()
    {
        return $this->action;
    }

    /**
     * Generated from protobuf field <code>.helmProxy.AuthAction action = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setAction($var)
    {
        GPBUtil::checkEnum($var, \MeshelmProxy\AuthAction::class);
        $this->action = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .helmProxy.AuthnRule rules = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getRules()
    {
        return $this->rules;
    }

    /**
     * Generated from protobuf field <code>repeated .helmProxy.AuthnRule rules = 2;</code>
     * @param \MeshelmProxy\AuthnRule[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setRules($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \MeshelmProxy\AuthnRule::class);
        $this->rules = $arr;

        return $this;
    }

}

