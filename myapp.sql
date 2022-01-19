-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 19 Jan 2022 pada 01.30
-- Versi server: 10.4.20-MariaDB
-- Versi PHP: 7.3.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `myapp`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `catalogues`
--

CREATE TABLE `catalogues` (
  `id` int(11) NOT NULL,
  `model_name` varchar(50) NOT NULL,
  `image_url` text NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `catalogues`
--

INSERT INTO `catalogues` (`id`, `model_name`, `image_url`, `description`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Models 50', 'https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.wikihow', 'simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,', '2022-01-18 19:13:52', '2022-01-18 19:47:23', NULL),
(3, 'Models 2', 'https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.wikihow', 'simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,', '2022-01-18 19:14:56', '2022-01-18 19:14:56', '2022-01-18 21:52:05'),
(4, 'Models 3', 'https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.wikihow', 'simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,', '2022-01-18 19:15:01', '2022-01-18 19:15:01', '2022-01-18 21:52:22'),
(5, 'Models 10', 'https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.wikihow', 'simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,', '2022-01-18 19:15:11', '2022-01-18 19:26:07', '2022-01-18 19:41:23'),
(6, 'Models 10', 'https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.wikihow', 'simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,', '2022-01-18 22:10:08', '2022-01-18 22:10:08', NULL),
(7, 'Models 11', 'https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.wikihow', 'simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,', '2022-01-18 22:10:14', '2022-01-18 22:10:14', NULL),
(8, 'Models 12', 'https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.wikihow', 'simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,', '2022-01-18 22:10:18', '2022-01-18 22:10:18', NULL),
(9, 'Models 9', 'https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.wikihow', 'simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,', '2022-01-18 22:10:22', '2022-01-18 22:12:15', '2022-01-18 22:12:36'),
(10, 'Models 13', 'https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.wikihow', 'simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages,', '2022-01-19 00:29:19', '2022-01-19 00:29:19', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `schedules`
--

CREATE TABLE `schedules` (
  `id` int(11) NOT NULL,
  `model_id` int(11) NOT NULL,
  `schedule_date` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `schedules`
--

INSERT INTO `schedules` (`id`, `model_id`, `schedule_date`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 3, '2012-01-05 00:00:00', '2022-01-18 19:55:27', '2022-01-18 21:44:03', '2022-01-18 21:52:05'),
(2, 7, '2012-01-05 00:00:00', '2022-01-18 20:00:31', '2022-01-18 22:27:01', '2022-01-18 22:27:21'),
(3, 1, '2012-01-06 00:00:00', '2022-01-18 22:12:44', '2022-01-18 22:12:44', NULL),
(4, 6, '2012-01-05 00:00:00', '2022-01-18 22:12:51', '2022-01-18 22:12:51', NULL),
(5, 1, '2012-01-05 00:00:00', '2022-01-18 22:13:01', '2022-01-18 22:13:01', NULL),
(6, 6, '2012-01-05 00:00:00', '2022-01-18 22:13:02', '2022-01-19 00:26:16', NULL);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `catalogues`
--
ALTER TABLE `catalogues`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `schedules`
--
ALTER TABLE `schedules`
  ADD PRIMARY KEY (`id`),
  ADD KEY `schedules_catalogues` (`model_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `catalogues`
--
ALTER TABLE `catalogues`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT untuk tabel `schedules`
--
ALTER TABLE `schedules`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `schedules`
--
ALTER TABLE `schedules`
  ADD CONSTRAINT `schedules_catalogues` FOREIGN KEY (`model_id`) REFERENCES `catalogues` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
