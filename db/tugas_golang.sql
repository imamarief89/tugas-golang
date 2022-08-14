-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 14 Agu 2022 pada 05.46
-- Versi server: 10.4.20-MariaDB
-- Versi PHP: 8.0.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tugas_golang`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `tbl_taks`
--

CREATE TABLE `tbl_taks` (
  `id` int(11) NOT NULL,
  `taks_name` text DEFAULT NULL,
  `assignor` varchar(255) DEFAULT NULL,
  `deadline` varchar(255) DEFAULT NULL,
  `finish` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `tbl_taks`
--

INSERT INTO `tbl_taks` (`id`, `taks_name`, `assignor`, `deadline`, `finish`) VALUES
(4, 'Tugas Golang pertama', 'ita', 'ggsg', 1),
(5, 'Tugas Golang pertama', 'sgsg', 'ssf', 1),
(7, 'imam lagi', 'ita', '2022-08-20', 1),
(11, 'Tugas Golang pertama', 'dssdg', '2022-08-19', 1),
(12, 'Tugas Golang pertama', 'sfaf', '2022-08-27', 1),
(14, 'ddsgds', 'dsfdsf', '2022-08-22', 1),
(15, 'ddsgds', 'ita', '2022-08-19', 1),
(16, 'fff', 'bb', '2022-08-26', 1),
(17, 'Tugas Golang r', 'ccbcb', '2022-08-20', 0);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `tbl_taks`
--
ALTER TABLE `tbl_taks`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `tbl_taks`
--
ALTER TABLE `tbl_taks`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
