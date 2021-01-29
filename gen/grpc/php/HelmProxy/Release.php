<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace HelmProxy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Release tracks the state of an app during the lifecycle of it's current deployment
 *
 * Generated from protobuf message <code>helmProxy.Release</code>
 */
class Release extends \Google\Protobuf\Internal\Message
{
    /**
     * version of the App. Iterates with each upgrade
     *
     * Generated from protobuf field <code>uint32 version = 1;</code>
     */
    private $version = 0;
    /**
     * config values
     *
     * Generated from protobuf field <code>.google.protobuf.Struct config = 2;</code>
     */
    private $config = null;
    /**
     * notes associated with the release
     *
     * Generated from protobuf field <code>string notes = 3;</code>
     */
    private $notes = '';
    /**
     * description of the release
     *
     * Generated from protobuf field <code>string description = 4;</code>
     */
    private $description = '';
    /**
     * status of the release
     *
     * Generated from protobuf field <code>string status = 5;</code>
     */
    private $status = '';
    /**
     * lifecycle timestamps related
     *
     * Generated from protobuf field <code>.helmProxy.Timestamps timestamps = 6;</code>
     */
    private $timestamps = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $version
     *           version of the App. Iterates with each upgrade
     *     @type \Google\Protobuf\Struct $config
     *           config values
     *     @type string $notes
     *           notes associated with the release
     *     @type string $description
     *           description of the release
     *     @type string $status
     *           status of the release
     *     @type \HelmProxy\Timestamps $timestamps
     *           lifecycle timestamps related
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * version of the App. Iterates with each upgrade
     *
     * Generated from protobuf field <code>uint32 version = 1;</code>
     * @return int
     */
    public function getVersion()
    {
        return $this->version;
    }

    /**
     * version of the App. Iterates with each upgrade
     *
     * Generated from protobuf field <code>uint32 version = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setVersion($var)
    {
        GPBUtil::checkUint32($var);
        $this->version = $var;

        return $this;
    }

    /**
     * config values
     *
     * Generated from protobuf field <code>.google.protobuf.Struct config = 2;</code>
     * @return \Google\Protobuf\Struct
     */
    public function getConfig()
    {
        return $this->config;
    }

    /**
     * config values
     *
     * Generated from protobuf field <code>.google.protobuf.Struct config = 2;</code>
     * @param \Google\Protobuf\Struct $var
     * @return $this
     */
    public function setConfig($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Struct::class);
        $this->config = $var;

        return $this;
    }

    /**
     * notes associated with the release
     *
     * Generated from protobuf field <code>string notes = 3;</code>
     * @return string
     */
    public function getNotes()
    {
        return $this->notes;
    }

    /**
     * notes associated with the release
     *
     * Generated from protobuf field <code>string notes = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setNotes($var)
    {
        GPBUtil::checkString($var, True);
        $this->notes = $var;

        return $this;
    }

    /**
     * description of the release
     *
     * Generated from protobuf field <code>string description = 4;</code>
     * @return string
     */
    public function getDescription()
    {
        return $this->description;
    }

    /**
     * description of the release
     *
     * Generated from protobuf field <code>string description = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setDescription($var)
    {
        GPBUtil::checkString($var, True);
        $this->description = $var;

        return $this;
    }

    /**
     * status of the release
     *
     * Generated from protobuf field <code>string status = 5;</code>
     * @return string
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * status of the release
     *
     * Generated from protobuf field <code>string status = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkString($var, True);
        $this->status = $var;

        return $this;
    }

    /**
     * lifecycle timestamps related
     *
     * Generated from protobuf field <code>.helmProxy.Timestamps timestamps = 6;</code>
     * @return \HelmProxy\Timestamps
     */
    public function getTimestamps()
    {
        return $this->timestamps;
    }

    /**
     * lifecycle timestamps related
     *
     * Generated from protobuf field <code>.helmProxy.Timestamps timestamps = 6;</code>
     * @param \HelmProxy\Timestamps $var
     * @return $this
     */
    public function setTimestamps($var)
    {
        GPBUtil::checkMessage($var, \HelmProxy\Timestamps::class);
        $this->timestamps = $var;

        return $this;
    }

}

