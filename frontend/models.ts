export enum SearchType {
  Artist = 'Artist',
  Label = 'Label',
}

export interface SearchResults<T = ArtistSearchResult | LabelSearchResult> {
  results: T[];
}

export interface ArtistSearchResult {
  id: string;
  gid: string;
  name: string;
  sortName: string
  beginDateYear: number | null;
  beginDateMonth: number | null;
  beginDateDay: number | null;
  endDateYear: number | null;
  endDateMonth: number | null;
  endDateDay: number | null;
  type: number | null;
  comment: string;
  areaName: string | null;
  score: number;
}

export interface LabelSearchResult {
  id: string;
  gid: string;
  name: string;
  beginDateYear: number | null;
  beginDateMonth: number | null;
  beginDateDay: number | null;
  endDateYear: number | null;
  endDateMonth: number | null;
  endDateDay: number | null;
  labelCode: number | null;
  type: number | null;
  comment: string;
  areaName: string;
  score: number;
}

export type SearchResultType<T extends SearchType> =
  T extends SearchType.Artist ? ArtistSearchResult : LabelSearchResult;
