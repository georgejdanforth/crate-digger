import { SearchResults, SearchResultType, SearchType } from "./models";


export async function search<T extends SearchType>(
  searchType: T,
  query: string
): Promise<SearchResultType<T>[]> {
  const url = `/api/search?entity=${searchType.toLowerCase()}&query=${query}`;
  // TODO: error handling
  const results: SearchResults<SearchResultType<T>> = await fetch(url)
    .then((response) => response.json());
  return results.results;
}
