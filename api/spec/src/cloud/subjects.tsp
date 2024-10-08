import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

namespace OpenMeterCloud;

@route("/api/v1/subjects")
@tag("Subjects")
interface Subjects {
  /**
   * List subjects.
   */
  @get
  @operationId("listSubjects")
  list(): Subject[] | OpenMeter.CommonErrors;

  /**
   * Get subject by ID or key.
   */
  @get
  @operationId("getSubject")
  get(@path subjectIdOrKey: string): Subject | OpenMeter.NotFoundError | OpenMeter.CommonErrors;

  /**
   * Upserts a subject. Creates or updates subject.
   *
   * If the subject doesn't exist, it will be created.
   * If the subject exists, it will be partially updated with the provided fields.
   */
  @post
  @operationId("listSubjects")
  upsert(@body subject: Subject[]): Subject[] | OpenMeter.CommonErrors;
}

/**
 * A subject is a unique identifier for a user or entity.
 */
@friendlyName("Subject")
@example(#{
  id: "01G65Z755AFWAKHE12NY0CQ9FH",
  key: "customer-id",
  displayName: "Customer Name",
  metadata: #{ hubspotId: "123456" },
  currentPeriodStart: DateTime.fromISO("2023-01-01T00:00:00Z"),
  currentPeriodEnd: DateTime.fromISO("2023-02-01T00:00:00Z"),
  stripeCustomerId: "cus_JMOlctsKV8",
})
model Subject {
  // Validator doesn't obey required for readOnly properties
  // See: https://github.com/stoplightio/spectral/issues/1274

  /**
   * A unique identifier for the subject.
   */
  @visibility("read")
  @example("01G65Z755AFWAKHE12NY0CQ9FH")
  id: ULID;

  /**
   * A unique, human-readable identifier for the subject.
   * Must consist only alphanumeric and underscore characters.
   */
  @example("customer-id")
  key: Key;

  /**
   * A human-readable display name for the subject.
   */
  @example("Customer Name")
  displayName?: string;

  @example(#{ hubspotId: "123456" })
  metadata?: Record<string>;

  @example(DateTime.fromISO("2023-01-01T00:00:00Z"))
  currentPeriodStart?: DateTime;

  @example(DateTime.fromISO("2023-02-01T00:00:00Z"))
  currentPeriodEnd?: DateTime;

  @example("cus_JMOlctsKV8")
  stripeCustomerId?: string;
}
