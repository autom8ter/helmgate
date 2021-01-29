<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace HelmProxy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Timestamps tracks timestamps related to a release
 *
 * Generated from protobuf message <code>helmProxy.Timestamps</code>
 */
class Timestamps extends \Google\Protobuf\Internal\Message
{
    /**
     * when the release was first deployed
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp created = 1;</code>
     */
    private $created = null;
    /**
     * when the release was last deployed
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp updated = 2;</code>
     */
    private $updated = null;
    /**
     * when the release was deleted
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp deleted = 3;</code>
     */
    private $deleted = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Google\Protobuf\Timestamp $created
     *           when the release was first deployed
     *     @type \Google\Protobuf\Timestamp $updated
     *           when the release was last deployed
     *     @type \Google\Protobuf\Timestamp $deleted
     *           when the release was deleted
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * when the release was first deployed
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp created = 1;</code>
     * @return \Google\Protobuf\Timestamp
     */
    public function getCreated()
    {
        return $this->created;
    }

    /**
     * when the release was first deployed
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp created = 1;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setCreated($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->created = $var;

        return $this;
    }

    /**
     * when the release was last deployed
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp updated = 2;</code>
     * @return \Google\Protobuf\Timestamp
     */
    public function getUpdated()
    {
        return $this->updated;
    }

    /**
     * when the release was last deployed
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp updated = 2;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setUpdated($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->updated = $var;

        return $this;
    }

    /**
     * when the release was deleted
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp deleted = 3;</code>
     * @return \Google\Protobuf\Timestamp
     */
    public function getDeleted()
    {
        return $this->deleted;
    }

    /**
     * when the release was deleted
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp deleted = 3;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setDeleted($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->deleted = $var;

        return $this;
    }

}

