import { useQuery } from '@tanstack/react-query'
import axios from 'axios'
import { useState } from 'react'
import { useParams } from 'react-router-dom'
import { IEpisode } from 'src/components/hero'
import { times } from 'lodash'

const episodesPerPage = 10

export const Anime = () => {
  const params = useParams()
  const [current, setEpisode] = useState<IEpisode | null>(null)
  const [currentPage, setCurrentPage] = useState(1)

  const handlePageChange = (page: number) => {
    setCurrentPage(page)
  }

  const { data: episodes = [] } = useQuery({
    enabled: !!params.id,
    queryKey: ['animes', params.id],
    queryFn: async ({ signal }) => {
      const { data } = await axios.get<IEpisode[]>(`/api/animes/${params.id}`, {
        signal,
      })

      return data
    },
    onSuccess: (result) => setEpisode(result[0]),
  })
  const totalPages = Math.ceil(episodes.length / episodesPerPage)

  const {
    data: currentSrc = '',
    isLoading,
    isFetching,
  } = useQuery({
    enabled: !!current?.id,
    queryKey: ['src', current?.link],
    queryFn: async ({ signal }) => {
      const { data } = await axios.get<string>(`/api/episode/${current?.id}/src`, {
        signal,
      })

      return data
    },
  })

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 to-gray-700 py-8 text-white">
      <div className="container mx-auto px-4">
        <div className="flex flex-col items-center">
          {isLoading || isFetching ? (
            <div className="h-72 w-full animate-pulse rounded-lg bg-gray-300"></div>
          ) : (
            <iframe
              src={currentSrc}
              title="Anime Episode"
              className="h-72 w-full rounded-lg bg-white shadow-lg"
              allowFullScreen
            ></iframe>
          )}
        </div>
        <div className="mt-8 flex flex-col flex-wrap justify-center gap-3">
          <div className="flex justify-center space-x-2">
            <button
              onClick={() => handlePageChange(currentPage - 1)}
              disabled={currentPage === 1}
              className="rounded-lg bg-gray-600 px-4 py-2 text-white disabled:opacity-50"
            >
              Previous
            </button>
            <span className="my-auto px-4 py-2">{`Page ${currentPage} of ${totalPages}`}</span>
            <button
              onClick={() => handlePageChange(currentPage + 1)}
              disabled={currentPage === totalPages}
              className="rounded-lg bg-gray-600 px-4 py-2 text-white disabled:opacity-50"
            >
              Next
            </button>
          </div>
          <div className="flex justify-center">
            {times(episodesPerPage).map((_, index) => {
              const episodeIndex = (currentPage - 1) * episodesPerPage + index
              if (episodeIndex >= episodes.length) return null
              const episode = episodes[episodeIndex]
              const isActive = episode.id === current?.id
              return (
                <button
                  key={episode.id}
                  className={`m-2 transform rounded-lg px-4 py-2 transition-transform duration-300 hover:scale-105 ${
                    isActive ? 'bg-gray-600 text-white' : 'bg-white text-black hover:bg-gray-300'
                  }`}
                  onClick={() => setEpisode(episode)}
                >
                  {episode.title}
                </button>
              )
            })}
          </div>
        </div>
      </div>
    </div>
  )

  return (
    <div className="min-h-screen flex-row flex-wrap bg-gradient-to-b from-stone-700 to-stone-500">
      {isLoading || isFetching ? (
        <div className="mx-auto h-[800px] min-w-full pt-20" />
      ) : (
        <iframe
          onClick={(e) => {
            e.stopPropagation()
          }}
          key={currentSrc}
          className="mx-auto min-w-full pt-20"
          height="800px"
          allowFullScreen
          src={currentSrc}
        />
      )}
      <div
        className={`space-2 flex max-h-96 w-full flex-1 flex-row flex-wrap justify-between gap-2 overflow-y-auto p-2`}
      >
        {episodes.map((episode) => (
          <span
            onClick={(e) => {
              e.preventDefault()
              e.stopPropagation()
              setEpisode(episode)
            }}
            key={episode.link}
            className={`text-md flex h-8 w-12 cursor-pointer flex-wrap justify-center border-2 border-gray-50 p-1 text-center font-semibold text-white`}
          >
            {episode.title}
          </span>
        ))}
      </div>
    </div>
  )
}
