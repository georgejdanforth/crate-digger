import { ReactElement } from 'react';

function composeDate(y: number | null, m: number | null, d: number | null): Date | null {
  return (
    y == null ? null :
    m == null ? new Date(y) :
    d == null ? new Date(y, m) : new Date(y, m, d)
  );
}

type MaybeHasDateRange = {
  beginDateYear: number | null,
  beginDateMonth: number | null,
  beginDateDay: number | null,
  endDateYear: number | null,
  endDateMonth: number | null,
  endDateDay: number | null,
}

interface DateRangeProps<T extends MaybeHasDateRange> {
  entity: T
}
function DateRange<T extends MaybeHasDateRange>({ entity }: DateRangeProps<T>): ReactElement<DateRangeProps<T>> {
  const beginDate = composeDate(entity.beginDateYear, entity.beginDateMonth, entity.beginDateDay);
  const endDate = composeDate(entity.endDateYear, entity.endDateMonth, entity.endDateDay);

  if (beginDate == null) {
    return <></>;
  }

  if (endDate == null) {
    return (
      <span>{beginDate.getFullYear()}—Present</span>
    );
  }

  return (
    <span>{beginDate.getFullYear()}—{endDate.getFullYear()}</span>
  );
}

export default DateRange;
