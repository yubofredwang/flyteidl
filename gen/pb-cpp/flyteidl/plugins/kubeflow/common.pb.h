// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/plugins/kubeflow/common.proto

#ifndef PROTOBUF_INCLUDED_flyteidl_2fplugins_2fkubeflow_2fcommon_2eproto
#define PROTOBUF_INCLUDED_flyteidl_2fplugins_2fkubeflow_2fcommon_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_flyteidl_2fplugins_2fkubeflow_2fcommon_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_flyteidl_2fplugins_2fkubeflow_2fcommon_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[1]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_flyteidl_2fplugins_2fkubeflow_2fcommon_2eproto();
namespace flyteidl {
namespace plugins {
namespace kubeflow {
class RunPolicy;
class RunPolicyDefaultTypeInternal;
extern RunPolicyDefaultTypeInternal _RunPolicy_default_instance_;
}  // namespace kubeflow
}  // namespace plugins
}  // namespace flyteidl
namespace google {
namespace protobuf {
template<> ::flyteidl::plugins::kubeflow::RunPolicy* Arena::CreateMaybeMessage<::flyteidl::plugins::kubeflow::RunPolicy>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace flyteidl {
namespace plugins {
namespace kubeflow {

enum RestartPolicy {
  RESTART_POLICY_ALWAYS = 0,
  RESTART_POLICY_ON_FAILURE = 1,
  RESTART_POLICY_NEVER = 2,
  RestartPolicy_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::min(),
  RestartPolicy_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::max()
};
bool RestartPolicy_IsValid(int value);
const RestartPolicy RestartPolicy_MIN = RESTART_POLICY_ALWAYS;
const RestartPolicy RestartPolicy_MAX = RESTART_POLICY_NEVER;
const int RestartPolicy_ARRAYSIZE = RestartPolicy_MAX + 1;

const ::google::protobuf::EnumDescriptor* RestartPolicy_descriptor();
inline const ::std::string& RestartPolicy_Name(RestartPolicy value) {
  return ::google::protobuf::internal::NameOfEnum(
    RestartPolicy_descriptor(), value);
}
inline bool RestartPolicy_Parse(
    const ::std::string& name, RestartPolicy* value) {
  return ::google::protobuf::internal::ParseNamedEnum<RestartPolicy>(
    RestartPolicy_descriptor(), name, value);
}
enum CleanPodPolicy {
  CLEANPOD_POLICY_NONE = 0,
  CLEANPOD_POLICY_RUNNING = 1,
  CLEANPOD_POLICY_ALL = 2,
  CleanPodPolicy_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::min(),
  CleanPodPolicy_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::google::protobuf::int32>::max()
};
bool CleanPodPolicy_IsValid(int value);
const CleanPodPolicy CleanPodPolicy_MIN = CLEANPOD_POLICY_NONE;
const CleanPodPolicy CleanPodPolicy_MAX = CLEANPOD_POLICY_ALL;
const int CleanPodPolicy_ARRAYSIZE = CleanPodPolicy_MAX + 1;

const ::google::protobuf::EnumDescriptor* CleanPodPolicy_descriptor();
inline const ::std::string& CleanPodPolicy_Name(CleanPodPolicy value) {
  return ::google::protobuf::internal::NameOfEnum(
    CleanPodPolicy_descriptor(), value);
}
inline bool CleanPodPolicy_Parse(
    const ::std::string& name, CleanPodPolicy* value) {
  return ::google::protobuf::internal::ParseNamedEnum<CleanPodPolicy>(
    CleanPodPolicy_descriptor(), name, value);
}
// ===================================================================

class RunPolicy final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:flyteidl.plugins.kubeflow.RunPolicy) */ {
 public:
  RunPolicy();
  virtual ~RunPolicy();

  RunPolicy(const RunPolicy& from);

  inline RunPolicy& operator=(const RunPolicy& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  RunPolicy(RunPolicy&& from) noexcept
    : RunPolicy() {
    *this = ::std::move(from);
  }

  inline RunPolicy& operator=(RunPolicy&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const RunPolicy& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const RunPolicy* internal_default_instance() {
    return reinterpret_cast<const RunPolicy*>(
               &_RunPolicy_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(RunPolicy* other);
  friend void swap(RunPolicy& a, RunPolicy& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline RunPolicy* New() const final {
    return CreateMaybeMessage<RunPolicy>(nullptr);
  }

  RunPolicy* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<RunPolicy>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const RunPolicy& from);
  void MergeFrom(const RunPolicy& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(RunPolicy* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // .flyteidl.plugins.kubeflow.CleanPodPolicy clean_pod_policy = 1;
  void clear_clean_pod_policy();
  static const int kCleanPodPolicyFieldNumber = 1;
  ::flyteidl::plugins::kubeflow::CleanPodPolicy clean_pod_policy() const;
  void set_clean_pod_policy(::flyteidl::plugins::kubeflow::CleanPodPolicy value);

  // int32 ttl_seconds_after_finished = 2;
  void clear_ttl_seconds_after_finished();
  static const int kTtlSecondsAfterFinishedFieldNumber = 2;
  ::google::protobuf::int32 ttl_seconds_after_finished() const;
  void set_ttl_seconds_after_finished(::google::protobuf::int32 value);

  // int32 active_deadline_seconds = 3;
  void clear_active_deadline_seconds();
  static const int kActiveDeadlineSecondsFieldNumber = 3;
  ::google::protobuf::int32 active_deadline_seconds() const;
  void set_active_deadline_seconds(::google::protobuf::int32 value);

  // int32 backoff_limit = 4;
  void clear_backoff_limit();
  static const int kBackoffLimitFieldNumber = 4;
  ::google::protobuf::int32 backoff_limit() const;
  void set_backoff_limit(::google::protobuf::int32 value);

  // @@protoc_insertion_point(class_scope:flyteidl.plugins.kubeflow.RunPolicy)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  int clean_pod_policy_;
  ::google::protobuf::int32 ttl_seconds_after_finished_;
  ::google::protobuf::int32 active_deadline_seconds_;
  ::google::protobuf::int32 backoff_limit_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_flyteidl_2fplugins_2fkubeflow_2fcommon_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// RunPolicy

// .flyteidl.plugins.kubeflow.CleanPodPolicy clean_pod_policy = 1;
inline void RunPolicy::clear_clean_pod_policy() {
  clean_pod_policy_ = 0;
}
inline ::flyteidl::plugins::kubeflow::CleanPodPolicy RunPolicy::clean_pod_policy() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.RunPolicy.clean_pod_policy)
  return static_cast< ::flyteidl::plugins::kubeflow::CleanPodPolicy >(clean_pod_policy_);
}
inline void RunPolicy::set_clean_pod_policy(::flyteidl::plugins::kubeflow::CleanPodPolicy value) {
  
  clean_pod_policy_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.kubeflow.RunPolicy.clean_pod_policy)
}

// int32 ttl_seconds_after_finished = 2;
inline void RunPolicy::clear_ttl_seconds_after_finished() {
  ttl_seconds_after_finished_ = 0;
}
inline ::google::protobuf::int32 RunPolicy::ttl_seconds_after_finished() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.RunPolicy.ttl_seconds_after_finished)
  return ttl_seconds_after_finished_;
}
inline void RunPolicy::set_ttl_seconds_after_finished(::google::protobuf::int32 value) {
  
  ttl_seconds_after_finished_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.kubeflow.RunPolicy.ttl_seconds_after_finished)
}

// int32 active_deadline_seconds = 3;
inline void RunPolicy::clear_active_deadline_seconds() {
  active_deadline_seconds_ = 0;
}
inline ::google::protobuf::int32 RunPolicy::active_deadline_seconds() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.RunPolicy.active_deadline_seconds)
  return active_deadline_seconds_;
}
inline void RunPolicy::set_active_deadline_seconds(::google::protobuf::int32 value) {
  
  active_deadline_seconds_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.kubeflow.RunPolicy.active_deadline_seconds)
}

// int32 backoff_limit = 4;
inline void RunPolicy::clear_backoff_limit() {
  backoff_limit_ = 0;
}
inline ::google::protobuf::int32 RunPolicy::backoff_limit() const {
  // @@protoc_insertion_point(field_get:flyteidl.plugins.kubeflow.RunPolicy.backoff_limit)
  return backoff_limit_;
}
inline void RunPolicy::set_backoff_limit(::google::protobuf::int32 value) {
  
  backoff_limit_ = value;
  // @@protoc_insertion_point(field_set:flyteidl.plugins.kubeflow.RunPolicy.backoff_limit)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace kubeflow
}  // namespace plugins
}  // namespace flyteidl

namespace google {
namespace protobuf {

template <> struct is_proto_enum< ::flyteidl::plugins::kubeflow::RestartPolicy> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::flyteidl::plugins::kubeflow::RestartPolicy>() {
  return ::flyteidl::plugins::kubeflow::RestartPolicy_descriptor();
}
template <> struct is_proto_enum< ::flyteidl::plugins::kubeflow::CleanPodPolicy> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::flyteidl::plugins::kubeflow::CleanPodPolicy>() {
  return ::flyteidl::plugins::kubeflow::CleanPodPolicy_descriptor();
}

}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_flyteidl_2fplugins_2fkubeflow_2fcommon_2eproto
