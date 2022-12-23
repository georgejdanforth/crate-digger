import type { NextPage } from 'next'
import Head from 'next/head'
import { FC, useEffect, useState } from 'react';
import { search } from '../backend';

import DateRange from '../components/DateRange';
import { ArtistSearchResult, LabelSearchResult, SearchResultType, SearchType } from '../models';
import { classes } from '../utils';

enum ViewState {
  Search = 'Search',
  Results = 'Results',
}

const Home: NextPage = () => {

  const [viewState, setViewState] = useState<ViewState>(ViewState.Search);
  const [searchType, setSearchType] = useState<SearchType>(SearchType.Artist);
  const [searchQuery, setSearchQuery] = useState<string>('');

  const childProps = { searchType, setSearchType, searchQuery, setSearchQuery };

  return (
    <>
      <Head>
        <title>Crate Digger — Search</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="h-screen">
        {(() => viewState === ViewState.Search
          ? <Search {...childProps} setViewState={setViewState} />
          : <Results {...childProps} setViewState={setViewState} />
        )()}
      </main>

      <footer/>
    </>
  )
}

interface SearchProps {
  searchType: SearchType;
  setSearchType: (searchType: SearchType) => void;
  searchQuery: string;
  setSearchQuery: (setSearchQuery: string) => void;
  setViewState: (viewState: ViewState) => void;
}
const Search: FC<SearchProps> = ({
  searchType,
  setSearchType,
  searchQuery,
  setSearchQuery,
  setViewState,
}) => (
  <form
    onSubmit={(event) => {
      event.preventDefault();
      setViewState(ViewState.Results)
    }}
    className="flex flex-grow flex-col justify-center items-center h-full gap-5"
  >
    <h1 className="font-bold text-4xl">Crate Digger</h1>
    <div className="flex flex-col gap-2 sm:w-2/3 md:w-2/3 lg:w-1/3">
      <SearchBar
        searchType={searchType}
        searchQuery={searchQuery}
        setSearchQuery={setSearchQuery}
      />
      <SearchTypeSelect searchType={searchType} setSearchType={setSearchType} />
    </div>
  </form>
);

interface SearchBarProps {
  searchType: SearchType;
  searchQuery: string;
  setSearchQuery: (searchQuery: string) => void;
}
const SearchBar: FC<SearchBarProps> = ({ searchType, searchQuery, setSearchQuery }) => (
  <div className="flex flex-row gap-1">
    <input
      type="text"
      value={searchQuery}
      onChange={(e) => setSearchQuery(e.target.value)}
      className="flex-grow outline-none px-2 py-2 text-xl border-2 rounded-lg focus:drop-shadow-md"
      placeholder={`Search ${searchType}s...`}
    />
  </div>
);

interface SearchTypeSelectProps {
  searchType: SearchType;
  setSearchType: (searchType: SearchType) => void;
}
const SearchTypeSelect: FC<SearchTypeSelectProps> = ({ searchType, setSearchType }) => (
  <div className="flex flex-rowg gap-3 w-full">
    {Object.values(SearchType).map((st) => (
      <button
        type="button"
        onClick={() => setSearchType(st)}
        className={classes(
          'flex-grow outline-none px-4 py-1 hover:bg-gray-100 rounded-lg',
          searchType === st ? 'bg-gray-200' : '',
        )}
        key={st}
      >
        {st}
      </button>
    ))}
  </div>
);

type ResultsState<T extends SearchType> =
  | { done: false }
  | { done: true, results: SearchResultType<T>[] };

interface ResultsProps {
  searchType: SearchType;
  searchQuery: string;
  setViewState: (viewState: ViewState) => void;
}
const Results: FC<ResultsProps> = ({
  searchType,
  searchQuery,
  setViewState,
}) => {
  const [state, setState] = useState<ResultsState<typeof searchType>>({ done: false });

  useEffect(() => {
    (async () => {
      const results = await search<typeof searchType>(searchType, searchQuery);
      setState({ done: true, results });
    })();
  }, []);

  return (
    <div className="flex flex-row justify-center items-center py-5">
      <div className="flex flex-col justify-center md:w-2/3 lg:w-1/3">
        <div className="p-4">
          <button
            className="rounded-lg bg-gray-100 p-2 mb-4 hover:bg-gray-200"
            onClick={() => {
              setViewState(ViewState.Search);
            }}
          >
            Go Back
          </button>
          <div>
            <b>{searchType}</b> results for <i>&ldquo;{searchQuery}&rdquo;</i>...
          </div>
          <hr />
        </div>
        {state.done && (
          <div className="flex flex-col gap-3">
            {state.results.map((result) => {
              switch (searchType) {
                case SearchType.Artist:
                  return <Artist artist={result as ArtistSearchResult} />
                case SearchType.Label:
                  return <Label label={result as LabelSearchResult} />
                default:
                  return <></>;
              }
            })}
          </div>
        )}
        {!state.done && <div>Loading...</div>}
      </div>
    </div>
  );
}

const Artist: FC<{ artist: ArtistSearchResult }> = ({ artist }) => (
  <a className="hover:cursor-pointer">
    <div className="group bg-white px-4 py-4 rounded-lg hover:drop-shadow-lg">
      <h2 className="text-lg font-bold group-hover:underline">{artist.name}</h2>
      <div className="text-gray-500">
        <i>
          {artist.areaName && <span>{artist.areaName}&nbsp;</span>}
          <DateRange entity={artist} />
        </i>
      </div>
      <p>{artist.comment}</p>
    </div>
  </a>
);

const Label: FC<{ label: LabelSearchResult }> = ({ label }) => (
  <a className="hover:cursor-pointer">
    <div className="group bg-white px-4 py-4 rounded-lg hover:drop-shadow-lg">
      <h2 className="text-lg font-bold group-hover:underline">{label.name}</h2>
      <div className="text-gray-500">
        <i>
          {label.areaName && <span>{label.areaName}&nbsp;</span>}
          <DateRange entity={label} />
        </i>
      </div>
      <p>{label.comment}</p>
    </div>
  </a>
);

export default Home
