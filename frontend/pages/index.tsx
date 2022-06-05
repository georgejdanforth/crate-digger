import type { NextPage } from 'next'
import Head from 'next/head'
import { FC, useState } from 'react';

import { classes } from '../utils';

enum SearchType {
  Artist = 'Artist',
  Label = 'Label',
}

const Home: NextPage = () => {

  const [searchType, setSearchType] = useState<SearchType>(SearchType.Artist);
  const [searchQuery, setSearchQuery] = useState<string>('');

  return (
    <>
      <Head>
        <title>Crate Digger — Search</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="h-screen">
        <div className="flex flex-grow flex-col justify-center items-center h-full gap-5">
          <h1 className="font-bold text-4xl">Crate Digger</h1>
          <div className="flex flex-col gap-2 sm:w-2/3 md:w-2/3 lg:w-1/3">
            <SearchBar
              searchType={searchType}
              searchQuery={searchQuery}
              setSearchQuery={setSearchQuery}
            />
            <SearchTypeSelect searchType={searchType} setSearchType={setSearchType} />
          </div>
        </div>
      </main>

      <footer/>
    </>
  )
}

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

export default Home
